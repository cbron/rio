package deploymentwrangler

import (
	"context"
	"fmt"

	"github.com/rancher/rio/modules/service/controllers/deploymentwrangler/populate"
	riov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	riov1controller "github.com/rancher/rio/pkg/generated/controllers/rio.cattle.io/v1"
	"github.com/rancher/rio/types"
	appsv1controller "github.com/rancher/wrangler-api/pkg/generated/controllers/apps/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
)

type deploymentWranglerHandler struct {
	//namespace               string
	deploymentWranglerCache riov1controller.DeploymentWranglerCache
	deploymentCache         appsv1controller.DeploymentCache
	appSelector             labels.Selector
}

func Register(ctx context.Context, rContext *types.Context) error {

	dwh := &deploymentWranglerHandler{
		deploymentCache:         rContext.Apps.Apps().V1().Deployment().Cache(),
		deploymentWranglerCache: rContext.Rio.Rio().V1().DeploymentWrangler().Cache(),
	}

	//rContext.Rio.Rio().V1().DeploymentWrangler().OnChange(ctx, "app", dwh.onDeploymentWranglerChange)
	//rContext.Apps.Apps ().V1().Deployment().OnChange(ctx, "app", dwh.onDeploymentChange)

	riov1controller.RegisterDeploymentWranglerGeneratingHandler(ctx,
		rContext.Rio.Rio().V1().DeploymentWrangler(),
		rContext.Apply.WithCacheTypes(
			rContext.Rio.Rio().V1().Service(),
			rContext.Apps.Apps().V1().Deployment()).
			WithRateLimiting(20), // is this right ?
		"DeploymentWranglerDeployed",
		"deploymentwrangler",
		dwh.generate,
		nil)

	return nil
}

func (dwh *deploymentWranglerHandler) generate(dw *riov1.DeploymentWrangler, status riov1.DeploymentWranglerStatus) ([]runtime.Object, riov1.DeploymentWranglerStatus, error) {
	fmt.Println("DW populate just fired", dw.Name)
	// check for deploy create special rio service, or update service to check itself ?

	// might already have 3 versions of the app. You join them up via spec.app and spec.version, but dw.name must equal deploy.name
	existingDeploy, err := dwh.deploymentCache.Get(dw.Namespace, dw.Name)

	if (err != nil && k8sErrors.IsNotFound(err)) || existingDeploy == nil {
		fmt.Println("deployment not found -> ", err)
		return nil, status, nil
	}
	if err != nil {
		return nil, status, err
	}
	service, err := populate.RioServiceForDeploymentWrangler(dw)
	if err != nil {
		return nil, status, err
	}
	return []runtime.Object{service}, status, nil
}

//
//func (dwh *deploymentWranglerHandler) onDeploymentWranglerChange(key string, dw *riov1.DeploymentWrangler) (*riov1.DeploymentWrangler, error) {
//	if dw == nil {
//		return nil, nil
//	}
//	//
//	//appName, _ := services.AppAndVersion(svc)
//	//revisions, err := h.serviceCache.GetByIndex(indexes.ServiceByApp, fmt.Sprintf("%s/%s", svc.Namespace, appName))
//	//if err != nil || len(revisions) == 0 {
//	//	return svc, err
//	//}
//
//	existingDeploy, err := dwh.deploymentCache.Get(dw.Namespace, dw.Name)
//	if err != nil {
//		fmt.Println("ERR not nil", err)
//		return dw, nil
//	}
//
//	//existingSvc, err := h.services.Cache().Get(svc.Namespace, appName)
//	//if err == nil {
//	//	ports := portsForService(revisions)
//	//	if !reflect.DeepEqual(existingSvc, ports) {
//	//		existingSvc.Spec.Ports = ports
//	//		if _, err := h.services.Update(existingSvc); err != nil {
//	//			return svc, err
//	//		}
//	//		return svc, nil
//	//	}
//	//	// Already Exists
//	//	return svc, nil
//	//}
//
//	//
//	//appService := newService(svc.Namespace, appName, revisions)
//	//_, err = h.services.Create(appService)
//	//if errors.IsAlreadyExists(err) {
//	//	// Already Exists
//	//	return svc, nil
//	//}
//
//	return svc, err
//}
