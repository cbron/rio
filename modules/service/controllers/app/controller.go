package app

import (
	"context"
	"fmt"
	"reflect"
	"sort"

	riov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	"github.com/rancher/rio/pkg/constructors"
	riov1controller "github.com/rancher/rio/pkg/generated/controllers/rio.cattle.io/v1"
	"github.com/rancher/rio/pkg/indexes"
	"github.com/rancher/rio/pkg/serviceports"
	"github.com/rancher/rio/pkg/services"
	"github.com/rancher/rio/types"
	corev1controller "github.com/rancher/wrangler-api/pkg/generated/controllers/core/v1"
	"github.com/rancher/wrangler/pkg/relatedresource"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const AppServiceLabel = "rio.cattle.io/app"

// The app controller creates and manages the app (un-versioned) k8s service
func Register(ctx context.Context, rContext *types.Context) error {
	appSelector, err := labels.Parse(AppServiceLabel)
	if err != nil {
		return err
	}

	h := &handler{
		dwCache:     rContext.Rio.Rio().V1().DeploymentWrangler().Cache(),
		sswCache:    rContext.Rio.Rio().V1().StatefulSetWrangler().Cache(),
		services:    rContext.Core.Core().V1().Service(),
		appSelector: appSelector,
	}

	rContext.Rio.Rio().V1().DeploymentWrangler().OnChange(ctx, "app", h.onDWServiceChange)
	rContext.Rio.Rio().V1().StatefulSetWrangler().OnChange(ctx, "app", h.onSSWServiceChange)
	rContext.Core.Core().V1().Service().OnChange(ctx, "app", h.onServiceChange)

	relatedresource.Watch(ctx, "app",
		resolveAppService,
		rContext.Core.Core().V1().Service(),
		rContext.Rio.Rio().V1().DeploymentWrangler(),
		rContext.Rio.Rio().V1().StatefulSetWrangler(),
	)

	return nil
}

// resolveAppService gets called on either workload update and enqueues its app
func resolveAppService(_, _ string, obj runtime.Object) ([]relatedresource.Key, error) {
	var app, ns string
	dw, ok := obj.(*riov1.DeploymentWrangler)
	if ok {
		app, _ = services.AppAndVersion(dw)
		ns = dw.Namespace
	} else {
		ssw, ok := obj.(*riov1.StatefulSetWrangler)
		if ok {
			app, _ = services.AppAndVersion(ssw)
			ns = dw.Namespace
		}
	}
	if app == "" {
		return nil, nil
	}
	return []relatedresource.Key{
		{
			Namespace: ns,
			Name:      app,
		},
	}, nil
}

type handler struct {
	dwCache     riov1controller.DeploymentWranglerCache
	sswCache    riov1controller.StatefulSetWranglerCache
	services    corev1controller.ServiceController
	appSelector labels.Selector
}

func (h *handler) onDWServiceChange(key string, dw *riov1.DeploymentWrangler) (*riov1.DeploymentWrangler, error) {
	if dw == nil {
		return nil, nil
	}
	appName, _ := services.AppAndVersion(dw)
	revisions, err := h.dwCache.GetByIndex(indexes.DWByApp, fmt.Sprintf("%s/%s", dw.Namespace, appName))
	if err != nil || len(revisions) == 0 {
		return dw, err
	}
	err = h.WorkloadChange(appName, dw.Namespace, riov1.DeploymentWranglerWorkloadSlice(revisions))
	if err != nil {
		return dw, err
	}
	return dw, nil
}

func (h *handler) onSSWServiceChange(key string, ssw *riov1.StatefulSetWrangler) (*riov1.StatefulSetWrangler, error) {
	if ssw == nil {
		return nil, nil
	}
	appName, _ := services.AppAndVersion(ssw)
	revisions, err := h.sswCache.GetByIndex(indexes.SSWByApp, fmt.Sprintf("%s/%s", ssw.Namespace, appName))
	if err != nil || len(revisions) == 0 {
		return ssw, err
	}
	err = h.WorkloadChange(appName, ssw.Namespace, riov1.StatefulSetWranglerWorkloadSlice(revisions))
	if err != nil {
		return ssw, err
	}
	return ssw, nil
}

// On workload change, gather ports from all revisions and set on app svc
func (h *handler) WorkloadChange(appName, ns string, revisions []riov1.Workload) error {
	existingSvc, err := h.services.Cache().Get(ns, appName)
	if err == nil {
		ports := portsForService(revisions)
		if !reflect.DeepEqual(existingSvc, ports) {
			existingSvc.Spec.Ports = ports
			if _, err := h.services.Update(existingSvc); err != nil {
				return err
			}
			return nil
		}
		// Already Exists
		return nil
	}

	appService := newService(ns, appName, revisions)
	_, err = h.services.Create(appService)
	if errors.IsAlreadyExists(err) {
		// Already Exists
		return nil
	}

	return err
}

// todo: Do we want app service to get deleted if DW is deleted ?
// onServiceChange checks if any rio workload exists for the app service, if not it removes it
func (h *handler) onServiceChange(key string, svc *corev1.Service) (*corev1.Service, error) {
	if svc == nil {
		return nil, nil
	}
	appName := svc.Labels[AppServiceLabel]
	if appName == "" {
		return svc, nil
	}

	dwInstances, err := h.dwCache.GetByIndex(indexes.DWByApp, fmt.Sprintf("%s/%s", svc.Namespace, appName))
	if err != nil {
		return svc, err
	}
	if len(dwInstances) > 0 {
		return svc, err
	}

	sswInstances, err := h.sswCache.GetByIndex(indexes.SSWByApp, fmt.Sprintf("%s/%s", svc.Namespace, appName))
	if err != nil {
		return svc, err
	}
	if len(sswInstances) > 0 {
		return svc, err
	}
	// no rio workload for this svc, remove it
	return svc, h.services.Delete(svc.Namespace, svc.Name, nil)
}

func newService(namespace, app string, serviceSet []riov1.Workload) *corev1.Service {
	ports := portsForService(serviceSet)
	return constructors.NewService(namespace, app, corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				AppServiceLabel: app,
			},
		},
		Spec: corev1.ServiceSpec{
			Ports: ports,
			Selector: map[string]string{
				"app": app,
			},
			Type: corev1.ServiceTypeClusterIP,
		},
	})
}

func portsForService(workloadSet []riov1.Workload) (result []corev1.ServicePort) {
	ports := map[string]corev1.ServicePort{}

	for _, rev := range workloadSet {
		for _, port := range serviceports.ServiceNamedPorts(rev) {
			ports[port.Name] = port
		}
	}

	for _, port := range ports {
		result = append(result, port)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	if len(result) == 0 {
		return []corev1.ServicePort{
			{
				Name:       "default",
				Protocol:   corev1.ProtocolTCP,
				TargetPort: intstr.FromInt(80),
				Port:       80,
			},
		}
	}

	return
}
