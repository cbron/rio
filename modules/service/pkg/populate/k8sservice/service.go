package k8sservice

import (
	riov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	"github.com/rancher/wrangler/pkg/objectset"
)

func Populate(shared riov1.Wrangler, os *objectset.ObjectSet) {
	serviceSelector(shared, os)
}
