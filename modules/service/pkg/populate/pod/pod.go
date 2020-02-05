package pod

//
//func Populate(service *riov1.Service, os *objectset.ObjectSet) (v1.PodTemplateSpec, error) {
//	pts := v1.PodTemplateSpec{
//		ObjectMeta: metav1.ObjectMeta{
//			Labels:      servicelabels.ServiceLabels(service),
//			Annotations: servicelabels.ServiceAnnotations(service),
//		},
//	}
//
//	podSpec := podSpec(service)
//	Roles(service, os)
//
//	PersistVolume(service, os)
//
//	pts.Spec = podSpec
//	return pts, nil
//}
//
////func Roles(service *riov1.Service, os *objectset.ObjectSet) {
////	if err := rbac.Populate(service, os); err != nil {
////		os.AddErr(err)
////		return
////	}
////}
//
//func PersistVolume(service *riov1.Service, os *objectset.ObjectSet) {
//	var volumes []riov1.Volume
//	for _, volume := range service.Spec.Volumes {
//		volumes = append(volumes, volume)
//	}
//
//	for _, c := range service.Spec.Containers {
//		for _, volume := range c.Volumes {
//			volumes = append(volumes, volume)
//		}
//	}
//
//	for _, v := range volumes {
//		if v.Persistent {
//			pv := constructors.NewPersistentVolumeClaim(service.Namespace, v.Name, v1.PersistentVolumeClaim{
//				Spec: v1.PersistentVolumeClaimSpec{
//					AccessModes: []v1.PersistentVolumeAccessMode{
//						v1.ReadWriteOnce,
//					},
//					Resources: v1.ResourceRequirements{
//						Requests: v1.ResourceList{
//							v1.ResourceStorage: resource.MustParse(constants.RegistryStorageSize),
//						},
//					},
//				},
//			})
//			os.Add(pv)
//		}
//	}
//}
