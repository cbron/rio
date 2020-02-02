package podcontrollers

//func daemonset(service *riov1.Service, cp *controllerParams, os *objectset.ObjectSet) {
//	ds := constructors.NewDaemonset(service.Namespace, service.Name, appsv1.DaemonSet{
//		ObjectMeta: metav1.ObjectMeta{
//			Labels:      cp.Labels,
//			PodAnnotations: cp.PodAnnotations,
//		},
//		Spec: appsv1.DaemonSetSpec{
//			Selector: &metav1.LabelSelector{
//				MatchLabels: cp.SelectorLabels,
//			},
//			Template: cp.PodTemplateSpec,
//			UpdateStrategy: appsv1.DaemonSetUpdateStrategy{
//				Type: appsv1.RollingUpdateDaemonSetStrategyType,
//				RollingUpdate: &appsv1.RollingUpdateDaemonSet{
//					MaxUnavailable: cp.Scale.MaxUnavailable,
//				},
//			},
//		},
//	})
//
//	os.Add(ds)
//}
//
