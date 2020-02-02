package k8sservice

import (
	"github.com/rancher/rio/modules/service/pkg/populate/labels"
	riov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	"github.com/rancher/rio/pkg/constructors"
	"github.com/rancher/rio/pkg/serviceports"
	"github.com/rancher/rio/pkg/services"
	"github.com/rancher/wrangler/pkg/objectset"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// Creates service for this version of app
func Populate(shared riov1.Wrangler, os *objectset.ObjectSet) {
	serviceSelector(shared, os)
}

func serviceSelector(w riov1.Wrangler, os *objectset.ObjectSet) {
	resourceLabels := labels.ResourceLabels(w)
	selectorLabels := labels.SelectorLabels(w)
	app, version := services.AppAndVersion(w) // todo: rename this services package, it references rio service
	svc := newServiceSelector(app+"-"+version, w.GetMeta().Namespace, v1.ServiceTypeClusterIP, resourceLabels, selectorLabels)
	if ports := serviceports.ServiceNamedPorts(w); len(ports) > 0 {
		svc.Spec.Ports = ports
	}
	os.Add(svc)
}

func newServiceSelector(name, namespace string, serviceType v1.ServiceType, resourceLabels, selectorLabels map[string]string) *v1.Service {
	return constructors.NewService(namespace, name, v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels: resourceLabels,
		},
		Spec: v1.ServiceSpec{
			Type:     serviceType,
			Selector: selectorLabels,
			Ports: []v1.ServicePort{
				{
					Name:       "default",
					Protocol:   v1.ProtocolTCP,
					TargetPort: intstr.FromInt(80),
					Port:       80,
				},
			},
		},
	})
}
