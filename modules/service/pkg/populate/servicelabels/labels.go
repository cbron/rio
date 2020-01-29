package servicelabels

import (
	v1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	"github.com/rancher/rio/pkg/services"
)

// Return annotations
func ServiceAnnotations(w v1.Wrangler) map[string]string {
	// user annotations will override ours
	return merge(annotations(w), w.GetMeta().Annotations)
}

// Return labels
func ServiceLabels(w v1.Wrangler) map[string]string {
	return merge(w.GetMeta().Labels, setLabels(w), SelectorLabels(w))
}

// Return labels needed for selecting
func SelectorLabels(w v1.Wrangler) map[string]string {
	app, version := services.AppAndVersion(w)
	return map[string]string{
		"app":     app,
		"version": version,
	}
}

func setLabels(w v1.Wrangler) map[string]string {
	return map[string]string{
		"rio.cattle.io/service": w.GetMeta().Name,
	}
}

func annotations(w v1.Wrangler) map[string]string {
	result := map[string]string{}
	if w.GetSpec().ServiceMesh != nil && !*w.GetSpec().ServiceMesh {
		result["rio.cattle.io/mesh"] = "false"
	} else {
		result["rio.cattle.io/mesh"] = "true"
	}
	return result
}

func merge(base map[string]string, overlay ...map[string]string) map[string]string {
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
		result = merge(merge(base, overlay[0]), overlay[1:]...)
	}

	return result
}
