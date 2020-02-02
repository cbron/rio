package podcontrollers

import (
	"github.com/rancher/rio/modules/service/pkg/populate/labels"
	"github.com/rancher/rio/pkg/constructors"
	"github.com/rancher/rio/pkg/services"
	"github.com/rancher/wrangler/pkg/objectset"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/equality"
)

func newDeployment(existing *appsv1.Deployment, cp *controllerParams, os *objectset.ObjectSet) bool {
	var deploy *appsv1.Deployment
	if existing == nil {
		// todo:  case where rio makes deployment from build, need to copy a base deploy not create from scratch. Need to change name.
		deploy = &appsv1.Deployment{
			//ObjectMeta: metav1.ObjectMeta{
			//	Labels:      cp.Labels,
			//	PodAnnotations: cp.PodAnnotations,
			//},
			//Spec: appsv1.DeploymentSpec{
			//	Replicas: cp.Scale.Scale,
			//	Selector: &metav1.LabelSelector{
			//		MatchLabels: cp.SelectorLabels,
			//	},
			//	Template: cp.PodTemplateSpec,
			//	Strategy: appsv1.DeploymentStrategy{
			//		Type: appsv1.RollingUpdateDeploymentStrategyType,
			//		RollingUpdate: &appsv1.RollingUpdateDeployment{
			//			MaxUnavailable: cp.Scale.MaxUnavailable,
			//			MaxSurge:       cp.Scale.MaxSurge,
			//		},
			//	},
			//},
		}
	} else {
		deploy = existing.DeepCopy()
		deploy.ObjectMeta = services.CleanMetadata(deploy.ObjectMeta)
	}

	if !allImagesSet(deploy.Spec.Template) {
		return false
	}
	mergeDeploymentLabels(deploy, cp)

	// make new deployment object, add to os, and determine if delete is necessary
	deploy = constructors.NewDeployment(deploy.Namespace, deploy.Name, *deploy)
	replace := false
	if existing != nil && !equality.Semantic.DeepEqual(existing.Spec.Selector, deploy.Spec.Selector) {
		replace = true // rio-made deploy needs to be created, delete existing deploy because selector is immutable
	}
	os.Add(deploy)
	return replace
}

func mergeDeploymentLabels(deploy *appsv1.Deployment, params *controllerParams) {
	// add deployment obj labels
	deploy.Labels = labels.Merge(deploy.Labels, params.ResourceLabels)

	// add deployment's pod selector labels
	deploy.Spec.Selector.MatchLabels = labels.Merge(deploy.Spec.Selector.MatchLabels, params.SelectorLabels)

	// add deployment's pod template labels and annotations
	deploy.Spec.Template.Labels = labels.Merge(deploy.Spec.Template.Labels, params.SelectorLabels)
	deploy.Spec.Template.Annotations = labels.Merge(deploy.Spec.Template.Annotations, params.MeshAnnotations)
}
