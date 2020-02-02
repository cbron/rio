package statefulsetwrangler

import (
	"context"
	"fmt"
	"github.com/rancher/rio/modules/service/controllers/deploymentwrangler"
	"github.com/rancher/rio/modules/service/pkg/populate/k8sservice"
	riov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	v1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	"github.com/rancher/rio/pkg/config"
	adminv1 "github.com/rancher/rio/pkg/generated/controllers/admin.rio.cattle.io/v1"
	riov1controller "github.com/rancher/rio/pkg/generated/controllers/rio.cattle.io/v1"
	"github.com/rancher/rio/types"
	appsv1controller "github.com/rancher/wrangler-api/pkg/generated/controllers/apps/v1"
	corev1controller "github.com/rancher/wrangler-api/pkg/generated/controllers/core/v1"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/objectset"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type statefulSetWranglerHandler struct {
	namespace          string
	statefulSetCache   appsv1controller.StatefulSetCache
	clusterDomainCache adminv1.ClusterDomainCache
	publicDomainCache  adminv1.PublicDomainCache
	configmaps         corev1controller.ConfigMapClient
}

func Register(ctx context.Context, rContext *types.Context) error {

	sswh := &statefulSetWranglerHandler{
		namespace:          rContext.Namespace,
		publicDomainCache:  rContext.Admin.Admin().V1().PublicDomain().Cache(),
		clusterDomainCache: rContext.Admin.Admin().V1().ClusterDomain().Cache(),
		statefulSetCache:   rContext.Apps.Apps().V1().StatefulSet().Cache(),
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
	existingSS, err := sswh.statefulSetCache.Get(ssw.Namespace, ssw.Name)
	if (err != nil && k8sErrors.IsNotFound(err)) || existingSS == nil {
		return nil, status, fmt.Errorf("statefulset \"%s\" not found", ssw.Name)
	}
	if err != nil {
		return nil, status, err
	}
	if ssw.Spec.WranglerSpec.Template {
		return nil, status, generic.ErrSkip
	}
	if err := sswh.ensureFeatures(ssw); err != nil {
		return nil, status, err
	}
	os := objectset.NewObjectSet()
	err = sswh.populate(ssw, os)
	if err != nil {
		return nil, status, err
	}
	return os.All(), status, nil
}

func (sswh *statefulSetWranglerHandler) populate(ssw *v1.StatefulSetWrangler, os *objectset.ObjectSet) error {
	k8sservice.Populate(ssw, os)
	//todo: add SS generator
	return nil
}

func (sswh *statefulSetWranglerHandler) ensureFeatures(ssw *riov1.StatefulSetWrangler) error {
	cm, err := sswh.configmaps.Get(sswh.namespace, config.ConfigName, metav1.GetOptions{})
	if err != nil {
		return err
	}
	err = deploymentwrangler.UpdateConfigForApp(cm, ssw)
	if err != nil {
		return err
	}
	if _, err := sswh.configmaps.Update(cm); err != nil {
		return err
	}
	return nil
}
