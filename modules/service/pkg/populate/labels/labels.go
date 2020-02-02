package labels

import (
	riov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	"github.com/rancher/rio/pkg/services"
)

// Return wrangler related labels with selectors and parent id attached. For k8s workloads and service, not pods.
func ResourceLabels(w riov1.Workload) map[string]string {
	return Merge(w.GetMeta().Labels, ParentLabels(w), SelectorLabels(w))
}

// Return labels needed for selecting this workload
func SelectorLabels(w riov1.Workload) map[string]string {
	app, version := services.AppAndVersion(w)
	return map[string]string{
		"app":     app,
		"version": version,
	}
}

func ParentLabels(w riov1.Workload) map[string]string {
	switch w.(type) {
	case riov1.DeploymentWrangler:
		return map[string]string{
			"rio.cattle.io/deploymentWrangler": w.GetMeta().Name,
		}
	case riov1.StatefulSetWrangler:
		return map[string]string{
			"rio.cattle.io/statefulSetWrangler": w.GetMeta().Name,
		}
	}
	return map[string]string{}
}

func MeshAnnotations(w riov1.Workload) map[string]string {
	result := map[string]string{}
	if w.GetSpec().ServiceMesh != nil && !*w.GetSpec().ServiceMesh {
		result["rio.cattle.io/mesh"] = "false"
	} else {
		result["rio.cattle.io/mesh"] = "true"
	}
	return result
}

func Merge(base map[string]string, overlay ...map[string]string) map[string]string {
	result := map[string]string{}
	for k, v := range base {
		result[k] = v
	}

	i := len(overlay)
	switch {
	case i == 1:
		for k, v := range overlay[0] {
			result[k] = v
		}
	case i > 1:
		result = Merge(Merge(base, overlay[0]), overlay[1:]...)
	}

	return result
}
