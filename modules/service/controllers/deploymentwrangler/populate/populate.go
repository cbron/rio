package populate

import (
	"errors"

	riov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	"github.com/rancher/rio/pkg/constants"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func RioServiceForDeploymentWrangler(dw *riov1.DeploymentWrangler) (*riov1.Service, error) {
	if dw.Spec.App == "" || dw.Spec.Version == "" {
		return nil, errors.New("DeploymentWrangler must specify app and version")
	}

	service := &riov1.Service{
		ObjectMeta: v1.ObjectMeta{
			Name:      dw.Spec.App + "-" + dw.Spec.Version,
			Namespace: dw.Namespace,
			Annotations: map[string]string{
				constants.DeploymentWranglerLabel: "true",
			},
		},
		Spec: riov1.ServiceSpec{
			App:             dw.Spec.App,
			Version:         dw.Spec.Version,
			Weight:          dw.Spec.Weight,
			Autoscale:       dw.Spec.Autoscale,
			RolloutDuration: dw.Spec.RolloutDuration,
		},
	}
	logrus.Infof("Generating rio service %s/%s from deployment wrangler", service.Namespace, service.Name)
	return service, nil
}
