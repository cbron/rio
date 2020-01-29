package populate

import (
	"github.com/rancher/rio/modules/service/pkg/populate/k8sservice"
	riov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	"github.com/rancher/wrangler/pkg/objectset"
)

func DeploymentWrangler(dw *riov1.DeploymentWrangler, os *objectset.ObjectSet) {
	k8sservice.Populate(dw, os)
}
