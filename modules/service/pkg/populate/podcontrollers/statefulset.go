package podcontrollers

//func statefulset(service *riov1.Service, cp *controllerParams, os *objectset.ObjectSet) {
//	appName, version := services.AppAndVersion(service)
//
//	ss := constructors.NewStatefulSet(service.Namespace, service.Name, appsv1.StatefulSet{
//		ObjectMeta: metav1.ObjectMeta{
//			Labels:      cp.Labels,
//			PodAnnotations: cp.PodAnnotations,
//		},
//		Spec: appsv1.StatefulSetSpec{
//			Replicas: nil,
//			Selector: &metav1.LabelSelector{
//				MatchLabels: cp.SelectorLabels,
//			},
//			Template:             cp.PodTemplateSpec,
//			VolumeClaimTemplates: volumeClaimTemplates(cp.VolumeTemplates),
//			ServiceName:          fmt.Sprintf("%s-%s", appName, version),
//			PodManagementPolicy:  appsv1.ParallelPodManagement,
//			UpdateStrategy: appsv1.StatefulSetUpdateStrategy{
//				Type: appsv1.RollingUpdateStatefulSetStrategyType,
//			},
//		},
//	})
//
//	os.Add(ss)
//}

//func volumeClaimTemplates(templates map[string]riov1.VolumeTemplate) (result []v1.PersistentVolumeClaim) {
//	var names []string
//	for name := range templates {
//		names = append(names, name)
//	}
//	sort.Strings(names)
//
//	for _, name := range names {
//		template := templates[name]
//		q := resource.NewQuantity(template.StorageRequest, resource.BinarySI)
//		result = append(result, v1.PersistentVolumeClaim{
//			ObjectMeta: metav1.ObjectMeta{
//				Name:        "vol-" + name,
//				Labels:      template.Labels,
//				PodAnnotations: template.PodAnnotations,
//			},
//			Spec: v1.PersistentVolumeClaimSpec{
//				AccessModes: template.AccessModes,
//				Resources: v1.ResourceRequirements{
//					Requests: v1.ResourceList{
//						v1.ResourceStorage: *q,
//					},
//				},
//				StorageClassName: &template.StorageClassName,
//				VolumeMode:       template.VolumeMode,
//			},
//		})
//	}
//
//	return
//}
