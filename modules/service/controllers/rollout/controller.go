package rollout

import (
	"context"
	"fmt"
	riov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	riov1controller "github.com/rancher/rio/pkg/generated/controllers/rio.cattle.io/v1"
	"github.com/rancher/rio/pkg/indexes"
	"github.com/rancher/rio/pkg/services"
	"github.com/rancher/rio/types"
	"sort"
	"time"
)

func Register(ctx context.Context, rContext *types.Context) error {
	rh := rolloutHandler{
		services:     rContext.Rio.Rio().V1().Service(),
		serviceCache: rContext.Rio.Rio().V1().Service().Cache(),
	}

	rContext.Rio.Rio().V1().Service().OnChange(ctx, "rollout", rh.sync)
	return nil
}

type rolloutHandler struct {
	services     riov1controller.ServiceController
	serviceCache riov1controller.ServiceCache
	client       riov1controller.ServiceClient
}

func (r rolloutHandler) sync(key string, obj *riov1.Service) (*riov1.Service, error) {
	// todo: remove all print statments
	fmt.Println("syncing...")
	return riov1controller.UpdateServiceDeepCopyOnChange(r.client, obj, r.rollout)
}

func (r rolloutHandler) rollout(svc *riov1.Service) (*riov1.Service, error) {

	// get all services
	appName, _ := services.AppAndVersion(svc)
	revisions, err := r.serviceCache.GetByIndex(indexes.ServiceByApp, fmt.Sprintf("%s/%s", svc.Namespace, appName))
	if err != nil || len(revisions) == 0 {
		return svc, err
	}

	// ensure revisions are in correct order so we get first rolloutConfig
	sort.Slice(revisions, func(i, j int) bool {
		return revisions[i].Spec.Version < revisions[j].Spec.Version
	})

	// When multiple services are initiated with no weight or computedWeight,
	// set initial ComputedWeights balanced evenly
	if !computedWeightsExist(revisions) {
		fmt.Println("setting initial computedWeights")
		var added int
		add := int(100.0 / float64(len(revisions)))
		for i, rev := range revisions {
			rev.Status.ComputedWeight = new(int)
			if i != len(revisions)-1 {
				fmt.Println("setting computed: ", svc.Name, add)
				*rev.Status.ComputedWeight = add
				added += add
			} else {
				fmt.Println("setting computed: ", svc.Name, 100-added)
				*rev.Status.ComputedWeight = 100 - added
			}
		}
	}

	// Check if ready and find a rolloutConfig
	ready := true
	var rolloutConfig *riov1.RolloutConfig
	for _, rev := range revisions {
		fmt.Printf("%v/%v Cond deployed: %v\n", rev.Namespace, rev.Name, riov1.ServiceConditionServiceDeployed.IsTrue(rev))
		// if this revision is not ready but has weight allocated break and return, can't scale until ready
		if riov1.ServiceConditionServiceDeployed.IsFalse(rev) && rev.Spec.Weight != nil && *rev.Spec.Weight > 0 {
			ready = false
			break
		}
		// grab first rolloutConfig found only
		if rolloutConfig == nil && rev.Spec.RolloutConfig != nil {
			rolloutConfig = rev.Spec.RolloutConfig
		}
	}
	// if services aren't ready or there are no rolloutConfigs found, return
	if !ready || !canRollout(rolloutConfig) {
		fmt.Println("Not Ready, RETURNING", svc.Name, ready, rolloutConfig)
		return svc, nil
	}

	for _, rev := range revisions {
		// loop through revisions and find one which has a spec.weight that's not yet met
		if rev.Spec.Weight == nil || (rev.Status.ComputedWeight != nil && *rev.Spec.Weight == *rev.Status.ComputedWeight) {
			fmt.Println("NOTHING TO DO, Returning", rev.Name)
			continue // this rev is already at desired weight, nothing to do
		}
		if rev.Status.ComputedWeight == nil {
			rev.Status.ComputedWeight = new(int)
		}
		fmt.Println("attempting to fix ", rev.Name, *rev.Status.ComputedWeight, *rev.Spec.Weight)
		observedWeight := *rev.Status.ComputedWeight

		// sleep in background and run again
		go func() {
			time.Sleep(rolloutConfig.Interval.Duration)
			r.services.Enqueue(rev.Namespace, rev.Name)
		}()

		weightToAdjust := *rev.Spec.Weight - observedWeight

		// calc weights and re-balance
		if incrementalRollout(rolloutConfig) {
			// if we can adjust less than entire increment, else whole increment
			if abs(weightToAdjust) < rolloutConfig.Increment {
				observedWeight += weightToAdjust
				magicSteal(revisions, -weightToAdjust)
			} else {
				rolloutAmount := rolloutConfig.Increment
				if weightToAdjust < 0 {
					rolloutAmount = -rolloutAmount
				}
				observedWeight += rolloutAmount
				magicSteal(revisions, -rolloutAmount)
			}
			*rev.Status.ComputedWeight = observedWeight
		} else {
			// immediate rollout
			*rev.Status.ComputedWeight += weightToAdjust
			magicSteal(revisions, -weightToAdjust)
		}
		break // only execute one revision at one sync call
	}
	return svc, nil
}

// canRollout confirms that we want to and are able to perform some rollout action
func canRollout(rc *riov1.RolloutConfig) bool {
	return rc != nil && rc.Pause != true
}

// incrementalRollout returns whether we want to perform intervaled rollout or immediate one
func incrementalRollout(rc *riov1.RolloutConfig) bool {
	return canRollout(rc) && rc.Increment != 0 && rc.Interval.Duration != 0
}

// Steal weight from other services and update them. Don't try to read it :)
func magicSteal(revisions []*riov1.Service, weightToAdjust int) {
	for _, rev := range revisions {
		if rev.Status.ComputedWeight == nil {
			continue // can't steal from nothing
		}
		specWeight := 0 // if spec weight never set, assume 0, meaning we can steal it all
		if rev.Spec.Weight != nil {
			specWeight = *rev.Spec.Weight
		}
		revAvailableWeight := *rev.Status.ComputedWeight - specWeight
		if revAvailableWeight == 0 {
			continue // if rev is at goal weight already, don't mess with it
		}

		// if rev's current weight - rev's goal weight * (neg)weightToAdjust is negative, we can steal
		// ex: 40 - 50 * -2 =  20   don't steal weight
		// ex: 50 - 40 * -2 = -20   steal weight
		if negative(revAvailableWeight, weightToAdjust) {
			// if amount this rev can adjust is greater than amount needed to adjust, do it all on this rev
			if abs(revAvailableWeight) > abs(weightToAdjust) {
				fmt.Printf("stealing %v from %v, which is now at %v\n", weightToAdjust, rev.Name, *rev.Status.ComputedWeight)
				*rev.Status.ComputedWeight += weightToAdjust
				weightToAdjust = 0
			} else { // steal just amount available
				fmt.Printf("stealing %v from %v, which is now at %v\n", weightToAdjust+revAvailableWeight, rev.Name, *rev.Status.ComputedWeight)
				weightToAdjust += revAvailableWeight
				*rev.Status.ComputedWeight = specWeight
			}
		}
		if weightToAdjust == 0 {
			fmt.Println("weightToAdjust is 0, breaking")
			break
		}
	}

	return
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func negative(a, b int) bool {
	return a*b < 0
}

func computedWeightsExist(revisions []*riov1.Service) bool {
	for _, rev := range revisions {
		if rev.Status.ComputedWeight != nil {
			return true
		}
	}
	return false
}
