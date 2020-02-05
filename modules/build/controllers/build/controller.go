package build

import (
	"context"
	"errors"
	"fmt"

	webhookv1controller "github.com/rancher/gitwatcher/pkg/generated/controllers/gitwatcher.cattle.io/v1"
	"github.com/rancher/rio/modules/build/controllers/service"
	v1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	"github.com/rancher/rio/pkg/constants"
	riov1controller "github.com/rancher/rio/pkg/generated/controllers/rio.cattle.io/v1"
	"github.com/rancher/rio/types"
	appsv1controller "github.com/rancher/wrangler-api/pkg/generated/controllers/apps/v1"
	"github.com/rancher/wrangler/pkg/condition"
	tektonv1alpha1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
)

type handler struct {
	systemNamespace  string
	deploymentClient appsv1controller.DeploymentClient
	deploymentCache  appsv1controller.DeploymentCache
	dwClient         riov1controller.DeploymentWranglerController
	sswClient        riov1controller.StatefulSetWranglerController
	stacks           riov1controller.StackController
	gitcommits       webhookv1controller.GitCommitController
}

func Register(ctx context.Context, rContext *types.Context) error {
	h := handler{
		systemNamespace:  rContext.Namespace,
		deploymentClient: rContext.Apps.Apps().V1().Deployment(),
		deploymentCache:  rContext.Apps.Apps().V1().Deployment().Cache(),
		dwClient:         rContext.Rio.Rio().V1().DeploymentWrangler(),
		sswClient:        rContext.Rio.Rio().V1().StatefulSetWrangler(),
		stacks:           rContext.Rio.Rio().V1().Stack(),
		gitcommits:       rContext.Webhook.Gitwatcher().V1().GitCommit(),
	}

	rContext.Build.Tekton().V1alpha1().TaskRun().OnChange(ctx, "build-service-update", h.updateWorkload)
	return nil
}

func (h handler) updateWorkload(key string, build *tektonv1alpha1.TaskRun) (*tektonv1alpha1.TaskRun, error) {
	if build == nil {
		return build, nil
	}

	var workload v1.Workload
	var dw *v1.DeploymentWrangler
	var ssw *v1.StatefulSetWrangler

	namespace, wName, wType, conName := build.Namespace, build.Labels[constants.WorkloadName], build.Labels[constants.WorkloadType], build.Labels[constants.ContainerLabel]
	if wType == constants.DeploymentWranglerType {
		var err error
		dw, err = h.dwClient.Cache().Get(namespace, wName)
		if err != nil {
			return build, nil
		}
		workload = dw
	} else if wType == constants.StatefulSetWranglerType {
		var err error
		ssw, err = h.sswClient.Cache().Get(namespace, wName)
		if err != nil {
			return build, nil
		}
		workload = ssw
	}
	if workload == nil {
		return build, nil
	}

	if workload.GetSpec().Template {
		return build, nil
	}

	state := ""
	if condition.Cond("Succeeded").IsFalse(build) {
		state = "failure"
	} else if condition.Cond("Succeeded").IsUnknown(build) {
		state = "in_progress"
	}

	if build.Labels[constants.GitCommitLabel] != "" {
		gitcommit, err := h.gitcommits.Cache().Get(build.Namespace, build.Labels[constants.GitCommitLabel])
		if err != nil {
			return build, err
		}
		gitcommit = gitcommit.DeepCopy()
		if gitcommit.Status.BuildStatus != state {
			gitcommit.Status.BuildStatus = state
			if _, err := h.gitcommits.Update(gitcommit); err != nil {
				return build, err
			}
		}
	}

	if condition.Cond("Succeeded").IsTrue(build) {
		imageName := ""
		for _, con := range workload.GetSpec().Containers {
			if con.Name == conName {
				imageName = service.PullImageName(namespace, conName, con.ImageBuild)
				break
			}
		}

		if wType == constants.DeploymentWranglerType {
			// First handle empty imageName case
			if imageName == "" {
				deepCopy := dw.DeepCopy()
				v1.ServiceConditionImageReady.SetError(deepCopy, "", fmt.Errorf("container name \"%s\" not found on deployment \"%s\"", conName, wName))
				_, err := h.dwClient.UpdateStatus(deepCopy)
				if err != nil {
					return build, err
				}
			}
			// If container found, update deployment
			deploy, err := h.deploymentCache.Get(namespace, wName)
			if err != nil {
				return build, err
			}
			for _, con := range deploy.Spec.Template.Spec.Containers {
				if con.Name == conName {
					deepCopy := deploy.DeepCopy()
					v1.ServiceConditionImageReady.SetError(deepCopy, "", nil)
					con.Image = imageName
					if _, err := h.deploymentClient.Update(deepCopy); err != nil {
						return build, err
					}
				}
			}
		} else if wType == constants.StatefulSetWranglerType {
			// todo: copy DW logic into here
		}

	} else if condition.Cond("Succeeded").IsFalse(build) {
		reason := condition.Cond("Succeeded").GetReason(build)
		message := condition.Cond("Succeeded").GetMessage(build)
		if wType == constants.DeploymentWranglerType {
			deepCopy := dw.DeepCopy()
			v1.ServiceConditionImageReady.SetError(deepCopy, reason, errors.New(message))
			_, err := h.dwClient.UpdateStatus(deepCopy)
			if err != nil {
				return build, err
			}
		} else if wType == constants.StatefulSetWranglerType {
			deepCopy := ssw.DeepCopy()
			v1.ServiceConditionImageReady.SetError(deepCopy, reason, errors.New(message))
			_, err := h.sswClient.UpdateStatus(deepCopy)
			if err != nil {
				return build, err
			}
		}
		return build, nil
	}

	return build, nil
}
