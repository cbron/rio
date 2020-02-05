package services

import (
	riov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func AppAndVersion(w riov1.Workload) (string, string) {
	app := w.GetSpec().App
	version := w.GetSpec().Version

	if app == "" {
		app = w.GetMeta().Name
	}
	if version == "" {
		if len(w.GetMeta().UID) < 8 {
			version = string(w.GetMeta().UID)
		} else {
			version = string(w.GetMeta().UID)[:8]
		}
	}
	return app, version
}

func CleanMetadata(meta v1.ObjectMeta) v1.ObjectMeta {
	meta.UID = ""
	meta.SelfLink = ""
	meta.ResourceVersion = ""
	meta.CreationTimestamp = metav1.Time{}
	meta.DeletionTimestamp = &metav1.Time{}
	return meta
}

func RootContainerName(w riov1.Workload) string {
	return w.GetMeta().Name
}

func containerIsValid(container riov1.Container) bool {
	//return container.Image != "" || container.ImageBuild != nil
	return container.ImageBuild != nil
}

// todo: cleanup
// Convert primary container to named container and return all of them
func ToNamedContainers(w riov1.Workload) (result []riov1.NamedContainer) {
	//if containerIsValid(w.GetSpec().Container) {
	//	result = append(result, riov1.NamedContainer{
	//		Name:      RootContainerName(w),
	//		Container: w.GetSpec().Container,
	//	})
	//}
	//result = append(result, w.GetSpec().Containers...)
	return w.GetSpec().Containers
}

func AutoscaleEnable(w riov1.Workload) bool {
	return w.GetSpec().Autoscale != nil && w.GetSpec().Autoscale.MinReplicas != nil && w.GetSpec().Autoscale.MaxReplicas != nil && *w.GetSpec().Autoscale.MinReplicas != *w.GetSpec().Autoscale.MaxReplicas
}
