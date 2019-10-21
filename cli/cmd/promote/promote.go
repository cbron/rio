package promote

import (
	"errors"
	"fmt"
	"time"

	"github.com/rancher/mapper"
	"github.com/rancher/rio/cli/pkg/clicontext"
	"github.com/rancher/rio/cli/pkg/types"
	riov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	"github.com/rancher/wrangler/pkg/kv"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type Promote struct {
	Increment int  `desc:"Increment value" default:"5"`
	Interval  int  `desc:"Interval value" default:"5"`
	Pause     bool `desc:"Whether to pause rollout or continue it. Default to false" default:"true"`
}

func (p *Promote) Run(ctx *clicontext.CLIContext) error {
	ctx.NoPrompt = true
	var allErrors []error
	namespace := ctx.GetSetNamespace()
	arg := ctx.CLI.Args()
	if !arg.Present() {
		return errors.New("at least one argument is needed")
	}
	// todo: allow just passing service name
	app, version := kv.Split(arg.First(), ":")
	if app == "" || version == "" {
		return errors.New("invalid app or version")
	}

	services, err := ctx.Rio.Services(namespace).List(metav1.ListOptions{})
	if err != nil {
		return err
	}

	var revisions []riov1.Service
	for _, rev := range services.Items {
		if rev.Spec.App == app {
			revisions = append(revisions, rev)
		}
	}
	if len(revisions) == 0 {
		return errors.New("no services found")
	}

	for _, rev := range revisions {
		err = ctx.UpdateResource(types.Resource{
			Namespace: rev.Namespace,
			Name:      rev.Name,
			Type:      types.ServiceType,
		}, func(obj runtime.Object) error {
			rev := obj.(*riov1.Service)
			if rev.Spec.Weight == nil {
				rev.Spec.Weight = new(int)
			}
			if rev.Spec.Version == version {
				*rev.Spec.Weight = 100
				rev.Spec.RolloutConfig = &riov1.RolloutConfig{
					Pause:     p.Pause,
					Increment: p.Increment,
					Interval: metav1.Duration{
						Duration: time.Duration(p.Interval) * time.Second,
					},
				}
				fmt.Printf("%s:%s promoted\n", rev.Spec.App, rev.Spec.Version)
			} else {
				*rev.Spec.Weight = 0
				rev.Spec.RolloutConfig = nil
			}
			return nil
		})
		allErrors = append(allErrors, err)
	}
	return mapper.NewErrors(allErrors...)
}
