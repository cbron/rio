package podcontrollers

import (
	"github.com/rancher/rio/modules/service/pkg/populate/labels"
	riov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	"github.com/rancher/wrangler/pkg/objectset"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

// Deployment adds any necessary objects for a DW and returns bool if a deploy deletion is necessary
func Deployment(dw *riov1.DeploymentWrangler, deploy *appsv1.Deployment, os *objectset.ObjectSet) (replaceDeployment bool) {
	params := newControllerParams(dw)
	replaceDeployment = newDeployment(deploy, params, os)
	return
}

// todo: fix up statefulSet
func StatefulSet(ssw riov1.StatefulSetWrangler, ss *appsv1.StatefulSet, os *objectset.ObjectSet) (replaceStatefulSet bool) {
	//params := newControllerParams(ssw)
	//replaceStatefulSet = statefulset(ss, params, os)
	return
}

type controllerParams struct {
	ResourceLabels  map[string]string
	ParentLabels    map[string]string
	MeshAnnotations map[string]string
	SelectorLabels  map[string]string
}

func newControllerParams(w riov1.Wrangler) *controllerParams {
	return &controllerParams{
		ResourceLabels:  labels.ResourceLabels(w),
		ParentLabels:    labels.ParentLabels(w),
		MeshAnnotations: labels.MeshAnnotations(w),
		SelectorLabels:  labels.SelectorLabels(w),
	}
}

func allImagesSet(podTemplate corev1.PodTemplateSpec) bool {
	for _, container := range podTemplate.Spec.Containers {
		if container.Image == "" {
			return false
		}
	}
	for _, container := range podTemplate.Spec.InitContainers {
		if container.Image == "" {
			return false
		}
	}
	return true
}
