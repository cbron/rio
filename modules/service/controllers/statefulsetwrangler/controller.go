package statefulsetwrangler

import (
	"context"

	"github.com/rancher/rio/modules/service/controllers/statefulsetwrangler/populate"

	riov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	adminv1 "github.com/rancher/rio/pkg/generated/controllers/admin.rio.cattle.io/v1"
	riov1controller "github.com/rancher/rio/pkg/generated/controllers/rio.cattle.io/v1"
	"github.com/rancher/rio/types"
	corev1controller "github.com/rancher/wrangler-api/pkg/generated/controllers/core/v1"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/objectset"
	"k8s.io/apimachinery/pkg/runtime"
)

type statefulSetWranglerHandler struct {
	namespace          string
	clusterDomainCache adminv1.ClusterDomainCache
	publicDomainCache  adminv1.PublicDomainCache
	configmaps         corev1controller.ConfigMapClient
}

func Register(ctx context.Context, rContext *types.Context) error {

	sswh := &statefulSetWranglerHandler{
		namespace:          rContext.Namespace,
		publicDomainCache:  rContext.Admin.Admin().V1().PublicDomain().Cache(),
		clusterDomainCache: rContext.Admin.Admin().V1().ClusterDomain().Cache(),
		configmaps:         rContext.Core.Core().V1().ConfigMap(),
	}

	riov1controller.RegisterStatefulSetWranglerGeneratingHandler(ctx,
		rContext.Rio.Rio().V1().StatefulSetWrangler(),
		rContext.Apply.WithCacheTypes(
			rContext.RBAC.Rbac().V1().Role(),
			rContext.RBAC.Rbac().V1().RoleBinding(),
			rContext.Apps.Apps().V1().StatefulSet(),
			rContext.Core.Core().V1().ServiceAccount(),
			rContext.Core.Core().V1().Secret()).
			WithRateLimiting(20), // is this right ?
		"StatefulSetWranglerDeployed",
		"statefulsetwrangler",
		sswh.generate,
		nil)

	return nil
}

func (sswh *statefulSetWranglerHandler) generate(ssw *riov1.StatefulSetWrangler, status riov1.StatefulSetWranglerStatus) ([]runtime.Object, riov1.StatefulSetWranglerStatus, error) {
	//if err := s.ensureFeatures(dw); err != nil {
	//	return nil, status, err
	//}
	if ssw.Spec.WranglerSpec.Template {
		return nil, status, generic.ErrSkip
	}
	os := objectset.NewObjectSet()
	populate.StatefulSetWrangler(ssw, os)
	return os.All(), status, nil
}

//func (dwh *deploymentWranglerHandler) ensureFeatures(dw *riov1.DeploymentWrangler) error {
//	cm, err := dwh.configmaps.Get(dwh.namespace, config.ConfigName, metav1.GetOptions{})
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
//	if _, err := dwh.configmaps.Update(cm); err != nil {
//		return err
//	}
//
//	return nil
//}
