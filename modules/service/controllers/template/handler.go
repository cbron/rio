package template

import (
	"context"
	"fmt"

	webhookv1controller "github.com/rancher/gitwatcher/pkg/generated/controllers/gitwatcher.cattle.io/v1"
	riov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	"github.com/rancher/rio/pkg/constants"
	riov1controller "github.com/rancher/rio/pkg/generated/controllers/rio.cattle.io/v1"
	"github.com/rancher/rio/pkg/indexes"
	"github.com/rancher/rio/pkg/services"
	"github.com/rancher/rio/types"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func Register(ctx context.Context, rContext *types.Context) error {
	h := &handler{
		dw:         rContext.Rio.Rio().V1().DeploymentWrangler(),
		gitcommits: rContext.Webhook.Gitwatcher().V1().GitCommit().Cache(),
	}

	riov1controller.RegisterDeploymentWranglerGeneratingHandler(ctx,
		rContext.Rio.Rio().V1().DeploymentWrangler(),
		rContext.Apply.
			WithCacheTypes(rContext.Rio.Rio().V1().DeploymentWrangler()).
			WithNoDelete(),
		"",
		"template",
		h.dwGenerate,
		nil)

	// todo: statefulset

	return nil
}

type handler struct {
	gitcommits webhookv1controller.GitCommitCache
	dw         riov1controller.DeploymentWranglerController
}

func (h *handler) dwGenerate(dw *riov1.DeploymentWrangler, status riov1.DeploymentWranglerStatus) ([]runtime.Object, riov1.DeploymentWranglerStatus, error) {
	objs, newStatus, err := h.generate(dw, status.WorkloadStatus)
	status = riov1.DeploymentWranglerStatus{
		WorkloadStatus: newStatus,
	}
	return objs, status, err
}

// generate takes a workload change and if its a template that needs to generates a new workload it creates it
func (h *handler) generate(w riov1.Workload, status riov1.WorkloadStatus) ([]runtime.Object, riov1.WorkloadStatus, error) {
	skip, err := h.skip(w)
	if err != nil {
		return nil, status, err
	}
	if skip {
		return nil, status, generic.ErrSkip
	}

	if err := h.cleanup(w); err != nil {
		return nil, status, err
	}

	name := status.ShouldGenerate
	app, _ := services.AppAndVersion(w)

	spec := w.GetSpec().DeepCopy()
	spec.Template = false
	spec.App = app
	spec.Version = ""
	setImageBuild(w, status, spec)
	setPullSecrets(spec)

	generatedFromPR, err := h.generatedFromPR(w)
	if err != nil {
		return nil, status, err
	}
	if !w.GetSpec().StageOnly && !generatedFromPR {
		svcs, err := h.services.Cache().GetByIndex(indexes.ServiceByApp, fmt.Sprintf("%s/%s", service.Namespace, app))
		if err != nil {
			return nil, status, err
		}
		duration := services.DefaultRolloutDuration
		if spec.RolloutDuration != nil {
			duration = spec.RolloutDuration.Duration
		}
		newWeight, rc, err := services.GenerateWeightAndRolloutConfig(service, svcs, 100, duration, false)
		if err != nil {
			return nil, status, err
		}
		spec.Weight = &newWeight
		spec.RolloutConfig = rc
		if err := h.scaleDownRevisions(service.Namespace, app, name, rc); err != nil {
			return nil, status, nil
		}
	}

	if status.ShouldClean[name] || status.GeneratedServices[name] {
		return nil, status, nil
	}

	logrus.Infof("Generating service %s/%s from template", service.Namespace, name)
	return []runtime.Object{
		&riov1.Service{
			ObjectMeta: v1.ObjectMeta{
				Name:      name,
				Namespace: service.Namespace,
				Annotations: map[string]string{
					constants.GitCommitLabel: last(service.Status.GitCommits),
				},
			},
			Spec: *spec,
		},
	}, status, nil
}

func (h *handler) generatedFromPR(w riov1.Workload) (bool, error) {
	if len(w.GetStatus().GitCommits) == 0 {
		return false, nil
	}

	gc, err := h.gitcommits.Get(w.GetMeta().Namespace, last(w.GetStatus().GitCommits))
	if err != nil {
		return false, err
	}

	return gc.Spec.PR != "", nil
}

func (h *handler) cleanup(w riov1.Workload) error {
	for shouldDelete := range w.GetStatus().ShouldClean {
		// todo: add ssw
		if err := h.dw.Delete(w.GetMeta().Namespace, shouldDelete, &metav1.DeleteOptions{}); err != nil && !errors.IsNotFound(err) {
			return err
		}
	}
	return nil
}

func (h *handler) scaleDownRevisions(namespace, name, excludedService string, rc *riov1.RolloutConfig) error {
	revisions, err := h.services.Cache().GetByIndex(indexes.ServiceByApp, fmt.Sprintf("%s/%s", namespace, name))
	if err != nil {
		return err
	}
	for _, revision := range revisions {
		if revision.Name == excludedService {
			continue
		}
		if revision.Spec.Template {
			continue
		}
		deepcopy := revision.DeepCopy()
		if deepcopy.Spec.Weight != nil && *deepcopy.Spec.Weight == 0 {
			continue
		}
		deepcopy.Spec.Weight = &[]int{0}[0]
		deepcopy.Spec.RolloutConfig = rc
		if _, err := h.services.Update(deepcopy); err != nil {
			return err
		}
		logrus.Infof("Scaling down service %s weight to 0", revision.Name)
	}
	return nil
}

func (h *handler) skip(w riov1.Workload) (bool, error) {
	if w.GetStatus().ShouldGenerate == "" {
		return true, nil
	}
	fromPR, err := h.generatedFromPR(w)
	if err != nil {
		return false, err
	}
	if fromPR {
		return false, nil
	}
	if !w.GetSpec().Template || len(w.GetStatus().ContainerRevision) == 0 {
		return true, nil
	}
	needed := 0
	has := 0
	// Revision empty is a template
	if w.GetSpec().ImageBuild != nil && w.GetSpec().ImageBuild.Revision == "" {
		needed++
	}
	for _, c := range w.GetSpec().Containers {
		if c.ImageBuild != nil && c.ImageBuild.Revision == "" {
			needed++
		}
	}

	for _, c := range w.GetStatus().ContainerRevision {
		if len(c.Commits) > 0 {
			has++
		}
	}
	return needed != has, nil
}

func setPullSecrets(ws *riov1.WorkloadSpec) {
	var imagePullSecrets []string

	//if ws.ImageBuild != nil && ws.ImageBuild.PushRegistrySecretName != "" {
	//	imagePullSecrets = append(imagePullSecrets, spec.ImageBuild.PushRegistrySecretName)
	//}

	for _, con := range ws.Containers {
		if con.ImageBuild != nil && con.ImageBuild.PushRegistrySecretName != "" {
			imagePullSecrets = append(imagePullSecrets, con.ImageBuild.PushRegistrySecretName)
		}
	}
}

func setImageBuild(w riov1.Workload, status riov1.WorkloadStatus, spec *riov1.WorkloadSpec) {
	//if w.GetSpec().ImageBuild != nil {
	//	spec.ImageBuild = w.GetSpec().ImageBuild
	//	spec.ImageBuild.Revision = last(status.ContainerRevision[services.RootContainerName(w)].Commits)
	//}

	for i := range spec.Containers {
		if w.GetSpec().Containers[i].ImageBuild != nil {
			spec.Containers[i].ImageBuild = w.GetSpec().Containers[i].ImageBuild
			spec.Containers[i].ImageBuild.Revision = last(status.ContainerRevision[spec.Containers[i].Name].Commits)
		}
	}
}

func last(a []string) string {
	if len(a) == 0 {
		return ""
	}
	return a[len(a)-1]
}
