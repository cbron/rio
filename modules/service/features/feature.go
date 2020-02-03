package features

import (
	"context"

	"github.com/rancher/rio/modules/service/controllers/app"
	"github.com/rancher/rio/modules/service/controllers/deploymentwrangler"
	"github.com/rancher/rio/modules/service/controllers/externalservice"
	"github.com/rancher/rio/modules/service/controllers/statefulsetwrangler"
	"github.com/rancher/rio/pkg/features"
	"github.com/rancher/rio/types"
)

func Register(ctx context.Context, rContext *types.Context) error {
	feature := &features.FeatureController{
		FeatureName: "service",
		FeatureSpec: features.FeatureSpec{
			Enabled:     true,
			Description: "Rio Service Based UX - required",
		},
		Controllers: []features.ControllerRegister{
			app.Register,
			externalservice.Register,
			//router.Register,
			deploymentwrangler.Register,
			statefulsetwrangler.Register,
			//globalrbac.Register,
			//servicestatus.Register,
			//rollout.Register,
			//template.Register,
		},
	}
	return feature.Register()
}
