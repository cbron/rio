package deploymentwrangler

//
//type serviceHandler struct {
//	namespace          string
//	clusterDomainCache adminv1.ClusterDomainCache
//	publicDomainCache  adminv1.PublicDomainCache
//	configmaps         corev1controller.ConfigMapClient
//}
//
//func Old(ctx context.Context, rContext *types.Context) error {
//
//	sh := &serviceHandler{
//		namespace:          rContext.Namespace,
//		publicDomainCache:  rContext.Admin.Admin().V1().PublicDomain().Cache(),
//		clusterDomainCache: rContext.Admin.Admin().V1().ClusterDomain().Cache(),
//		configmaps:         rContext.Core.Core().V1().ConfigMap(),
//	}
//
//	riov1controller.RegisterServiceGeneratingHandler(ctx,
//		rContext.Rio.Rio().V1().Service(),
//		rContext.Apply.WithCacheTypes(
//			rContext.RBAC.Rbac().V1().Role(),
//			rContext.RBAC.Rbac().V1().RoleBinding(),
//			rContext.Apps.Apps().V1().Deployment(),
//			rContext.Apps.Apps().V1().DaemonSet(),
//			rContext.Core.Core().V1().ServiceAccount(),
//			rContext.Core.Core().V1().Service(),
//			rContext.Core.Core().V1().Secret(),
//			rContext.Core.Core().V1().PersistentVolumeClaim()).
//			WithInjectorName("mesh").
//			WithRateLimiting(20),
//		"ServiceDeployed",
//		"service",
//		sh.populate,
//		nil)
//
//	return nil
//}
//
//func (s *serviceHandler) populate(service *riov1.Service, status riov1.ServiceStatus) ([]runtime.Object, riov1.ServiceStatus, error) {
//	if err := s.ensureFeatures(service); err != nil {
//		return nil, status, err
//	}
//
//	if service.Spec.Template {
//		return nil, status, generic.ErrSkip
//	}
//
//	os := objectset.NewObjectSet()
//	if err := populate.Service(service, os); err != nil {
//		return nil, status, err
//	}
//
//	return os.All(), status, nil
//}
//
//func (s *serviceHandler) ensureFeatures(service *riov1.Service) error {
//	cm, err := s.configmaps.Get(s.namespace, config.ConfigName, metav1.GetOptions{})
//	if err != nil {
//		return err
//	}
//
//	conf, err := config.FromConfigMap(cm)
//	if err != nil {
//		return err
//	}
//
//	t := true
//	if services.AutoscaleEnable(service) && arch.IsAmd64() {
//		if conf.Features == nil {
//			conf.Features = map[string]config.FeatureConfig{}
//		}
//		f := conf.Features["autoscaling"]
//		f.Enabled = &t
//		conf.Features["autoscaling"] = f
//	}
//
//	for _, con := range services.ToNamedContainers(service) {
//		if con.ImageBuild != nil && con.ImageBuild.Repo != "" && arch.IsAmd64() {
//			if conf.Features == nil {
//				conf.Features = map[string]config.FeatureConfig{}
//			}
//			f := conf.Features["build"]
//			f.Enabled = &t
//			conf.Features["build"] = f
//			break
//		}
//	}
//
//	cm, err = config.SetConfig(cm, conf)
//	if err != nil {
//		return err
//	}
//
//	if _, err := s.configmaps.Update(cm); err != nil {
//		return err
//	}
//
//	return nil
//}
