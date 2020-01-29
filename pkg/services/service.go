package services

import (
	riov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
)

func AppAndVersion(w riov1.Wrangler) (string, string) {
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

//func RootContainerName(w riov1.Wrangler) string {
//	return w.GetMeta().Name
//}
//
//func containerIsValid(container riov1.Container) bool {
//	return container.Image != "" || container.ImageBuild != nil
//}
//
//// Convert non-named container to named container using name
//func ToNamedContainers(w riov1.Wrangler) (result []riov1.NamedContainer) {
//	if containerIsValid(w.GetSpec().Container) {
//		result = append(result, riov1.NamedContainer{
//			Name:      RootContainerName(w),
//			Container: w.GetSpec().Container,
//		})
//	}
//
//	result = append(result, w.GetSpec().Sidecars...)
//	return
//}

func AutoscaleEnable(w riov1.Wrangler) bool {
	return w.GetSpec().Autoscale != nil && w.GetSpec().Autoscale.MinReplicas != nil && w.GetSpec().Autoscale.MaxReplicas != nil && *w.GetSpec().Autoscale.MinReplicas != *w.GetSpec().Autoscale.MaxReplicas
}
