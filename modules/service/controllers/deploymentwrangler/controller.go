package deploymentwrangler

import (
	"context"

	"github.com/rancher/rio/modules/service/pkg/populate/k8sservice"

	riov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	adminv1 "github.com/rancher/rio/pkg/generated/controllers/admin.rio.cattle.io/v1"
	riov1controller "github.com/rancher/rio/pkg/generated/controllers/rio.cattle.io/v1"
	"github.com/rancher/rio/types"
	corev1controller "github.com/rancher/wrangler-api/pkg/generated/controllers/core/v1"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/objectset"
	"k8s.io/apimachinery/pkg/runtime"
)

type deploymentWranglerHandler struct {
	namespace string
	//deploymentWranglerCache riov1controller.DeploymentWranglerCache
	//deploymentCache         appsv1controller.DeploymentCache
	clusterDomainCache adminv1.ClusterDomainCache
	publicDomainCache  adminv1.PublicDomainCache
	configmaps         corev1controller.ConfigMapClient
}

func Register(ctx context.Context, rContext *types.Context) error {

	dwh := &deploymentWranglerHandler{
		//deploymentCache:         rContext.Apps.Apps().V1().Deployment().Cache(),
		//deploymentWranglerCache: rContext.Rio.Rio().V1().DeploymentWrangler().Cache(),
		namespace:          rContext.Namespace,
		publicDomainCache:  rContext.Admin.Admin().V1().PublicDomain().Cache(),
		clusterDomainCache: rContext.Admin.Admin().V1().ClusterDomain().Cache(),
		configmaps:         rContext.Core.Core().V1().ConfigMap(),
	}

	riov1controller.RegisterDeploymentWranglerGeneratingHandler(ctx,
		rContext.Rio.Rio().V1().DeploymentWrangler(),
		rContext.Apply.WithCacheTypes(
			rContext.RBAC.Rbac().V1().Role(),
			rContext.RBAC.Rbac().V1().RoleBinding(),
			rContext.Apps.Apps().V1().Deployment(),
			rContext.Core.Core().V1().Service(),
			rContext.Core.Core().V1().ServiceAccount(),
			rContext.Core.Core().V1().Secret()).
			WithRateLimiting(20), // is this right ?
		"DeploymentWranglerDeployed",
		"deploymentwrangler",
		dwh.generate,
		nil)

	return nil
}

func (dwh *deploymentWranglerHandler) generate(dw *riov1.DeploymentWrangler, status riov1.DeploymentWranglerStatus) ([]runtime.Object, riov1.DeploymentWranglerStatus, error) {
	//if err := s.ensureFeatures(dw); err != nil {
	//	return nil, status, err
	//}
	if dw.Spec.WranglerSpec.Template {
		return nil, status, generic.ErrSkip
	}
	os := objectset.NewObjectSet()
	populate(dw, os)
	return os.All(), status, nil

	//existingDeploy, err := dwh.deploymentCache.Get(dw.Namespace, dw.Name)
	//
	//if err := populate.Service(service, os); err != nil {
	//	return nil, status, err
	//}
	//
	//if (err != nil && k8sErrors.IsNotFound(err)) || existingDeploy == nil {
	//	fmt.Println("deployment not found -> ", err)
	//	return nil, status, nil
	//}
	//if err != nil {
	//	return nil, status, err
	//}
	//service, err := populate.RioServiceForDeploymentWrangler(dw)
	//if err != nil {
	//	return nil, status, err
	//}
	//return []runtime.Object{service}, status, nil
}

func populate(dw *riov1.DeploymentWrangler, os *objectset.ObjectSet) {
	k8sservice.Populate(dw, os)
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
