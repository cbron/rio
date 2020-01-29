package pod

//
//var (
//	f             = false
//	t             = true
//	defaultCPU    = resource.MustParse("50m")
//	defaultMemory = resource.MustParse("64Mi")
//)
//
//func podSpec(service *riov1.Service) v1.PodSpec {
//	podSpec := v1.PodSpec{
//		DNSConfig:          podDNS(service),
//		HostAliases:        service.Spec.HostAliases,
//		Hostname:           service.Spec.Hostname,
//		HostNetwork:        service.Spec.HostNetwork,
//		EnableServiceLinks: &f,
//		Containers:         containers(service, false),
//		InitContainers:     containers(service, true),
//		Volumes:            volumes(service),
//		Affinity:           service.Spec.Affinity,
//		ImagePullSecrets:   pullSecrets(service.Spec.ImagePullSecrets),
//	}
//
//	serviceAccountName := rbac.ServiceAccountName(service)
//	if serviceAccountName != "" {
//		podSpec.ServiceAccountName = serviceAccountName
//		podSpec.AutomountServiceAccountToken = nil
//	}
//
//	if service.Spec.DNS != nil {
//		podSpec.DNSPolicy = service.Spec.DNS.Policy
//	}
//
//	if podSpec.ServiceAccountName == "" {
//		podSpec.AutomountServiceAccountToken = &f
//	} else {
//		podSpec.AutomountServiceAccountToken = &t
//	}
//
//	return podSpec
//}
//
//func pullSecrets(names []string) (result []v1.LocalObjectReference) {
//	for _, name := range names {
//		result = append(result, v1.LocalObjectReference{
//			Name: name,
//		})
//	}
//	return
//}
//
//func podDNS(service *riov1.Service) *v1.PodDNSConfig {
//	if service.Spec.DNS == nil {
//		return nil
//	}
//
//	if len(service.Spec.DNS.Options) == 0 &&
//		len(service.Spec.DNS.Nameservers) == 0 &&
//		len(service.Spec.DNS.Searches) == 0 {
//		return nil
//	}
//
//	var options []v1.PodDNSConfigOption
//	for _, opt := range service.Spec.DNS.Options {
//		options = append(options, v1.PodDNSConfigOption{
//			Name:  opt.Name,
//			Value: opt.Value,
//		})
//	}
//	return &v1.PodDNSConfig{
//		Options:     options,
//		Nameservers: service.Spec.DNS.Nameservers,
//		Searches:    service.Spec.DNS.Searches,
//	}
//}
