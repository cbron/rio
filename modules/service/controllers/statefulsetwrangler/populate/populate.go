package populate

import (
	"github.com/rancher/rio/modules/service/pkg/populate/k8sservice"
	v1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	"github.com/rancher/wrangler/pkg/objectset"
)

func StatefulSetWrangler(ssw *v1.StatefulSetWrangler, os *objectset.ObjectSet) {
	k8sservice.Populate(ssw, os)
}
