package deploymentwrangler

import (
	"context"
	"fmt"
	"github.com/rancher/rio/modules/service/pkg/populate/k8sservice"
	"github.com/rancher/rio/modules/service/pkg/populate/podcontrollers"
	riov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	"github.com/rancher/rio/pkg/arch"
	"github.com/rancher/rio/pkg/config"
	adminv1 "github.com/rancher/rio/pkg/generated/controllers/admin.rio.cattle.io/v1"
	riov1controller "github.com/rancher/rio/pkg/generated/controllers/rio.cattle.io/v1"
	"github.com/rancher/rio/pkg/services"
	"github.com/rancher/rio/types"
	appsv1controller "github.com/rancher/wrangler-api/pkg/generated/controllers/apps/v1"
	corev1controller "github.com/rancher/wrangler-api/pkg/generated/controllers/core/v1"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/objectset"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type deploymentWranglerHandler struct {
	namespace          string
	deploymentClient   appsv1controller.DeploymentClient
	deploymentCache    appsv1controller.DeploymentCache
	clusterDomainCache adminv1.ClusterDomainCache
	publicDomainCache  adminv1.PublicDomainCache
	configmaps         corev1controller.ConfigMapClient
}

func Register(ctx context.Context, rContext *types.Context) error {

	dwh := &deploymentWranglerHandler{
		deploymentCache:    rContext.Apps.Apps().V1().Deployment().Cache(),
		deploymentClient:   rContext.Apps.Apps().V1().Deployment(),
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
			WithRateLimiting(20),
		"DeploymentWranglerDeployed",
		"deploymentwrangler",
		dwh.generate,
		nil)
	return nil
}

// generate sets up a k8s deployment and svc for matching deployment-wrangler.
// Matches both on object name. Requires matching deployment only, will create svc if non-existing.
func (dwh *deploymentWranglerHandler) generate(dw *riov1.DeploymentWrangler, status riov1.DeploymentWranglerStatus) ([]runtime.Object, riov1.DeploymentWranglerStatus, error) {
	existingDeploy, err := dwh.deploymentCache.Get(dw.Namespace, dw.Name)
	if (err != nil && k8sErrors.IsNotFound(err)) || existingDeploy == nil {
		return nil, status, fmt.Errorf("deployment \"%s\" not found", dw.Name)
	}
	if err != nil {
		return nil, status, err
	}
	if dw.Spec.WorkloadSpec.Template {
		return nil, status, generic.ErrSkip
	}
	if err := dwh.ensureFeatures(dw); err != nil {
		return nil, status, err
	}
	os := objectset.NewObjectSet()
	err = dwh.populate(dw, existingDeploy, os)
	if err != nil {
		return nil, status, err
	}
	return os.All(), status, nil
}

func (dwh *deploymentWranglerHandler) populate(dw *riov1.DeploymentWrangler, existing *appsv1.Deployment, os *objectset.ObjectSet) error {
	k8sservice.Populate(dw, os)
	_ = podcontrollers.Deployment(dw, existing, os)
	//if replaceDeployment == true {
	//	fmt.Printf("DeploymentWrangler \"%s\" found matching deployment, recreating...\n", dw.Name)
	//	err := dwh.deploymentClient.Delete(existing.Namespace, existing.Name, &metav1.DeleteOptions{})
	//	if err != nil {
	//		return err
	//	}
	//}
	return nil
}

func (dwh *deploymentWranglerHandler) ensureFeatures(dw *riov1.DeploymentWrangler) error {
	cm, err := dwh.configmaps.Get(dwh.namespace, config.ConfigName, metav1.GetOptions{})
	if err != nil {
		return err
	}
	err = UpdateConfigForApp(cm, dw)
	if err != nil {
		return err
	}
	if _, err := dwh.configmaps.Update(cm); err != nil {
		return err
	}
	return nil
}

func UpdateConfigForApp(cm *v1.ConfigMap, w riov1.Workload) error {
	conf, err := config.FromConfigMap(cm)
	if err != nil {
		return err
	}

	t := true
	if services.AutoscaleEnable(w) && arch.IsAmd64() {
		if conf.Features == nil {
			conf.Features = map[string]config.FeatureConfig{}
		}
		f := conf.Features["autoscaling"]
		f.Enabled = &t
		conf.Features["autoscaling"] = f
	}

	if w.GetSpec().ImageBuild != nil && w.GetSpec().ImageBuild.Repo != "" && arch.IsAmd64() {
		if conf.Features == nil {
			conf.Features = map[string]config.FeatureConfig{}
		}
		f := conf.Features["build"]
		f.Enabled = &t
		conf.Features["build"] = f
	}

	cm, err = config.SetConfig(cm, conf)
	if err != nil {
		return err
	}
	return nil
}
