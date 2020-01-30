package feature

import (
	"context"

	"github.com/rancher/rio/pkg/features"
	"github.com/rancher/rio/pkg/stack"
	"github.com/rancher/rio/types"
)

func Register(ctx context.Context, rContext *types.Context) error {
	apply := rContext.Apply.WithCacheTypes(
		rContext.Rio.Rio().V1().DeploymentWrangler(),
		rContext.Rio.Rio().V1().StatefulSetWrangler(),
	)
	feature := &features.FeatureController{
		FeatureName: "autoscaling",
		FeatureSpec: features.FeatureSpec{
			Description: "Auto-scaling services based on in-flight requests",
		},
		SystemStacks: []*stack.SystemStack{
			stack.NewSystemStack(apply, rContext.Admin.Admin().V1().SystemStack(), rContext.Namespace, "rio-autoscaler"),
		},
		FixedAnswers: map[string]string{
			"NAMESPACE": rContext.Namespace,
		},
	}
	return feature.Register()
}
