package gitcommit

import (
	"context"
	webhookv1 "github.com/rancher/gitwatcher/pkg/apis/gitwatcher.cattle.io/v1"
	webhookv1controller "github.com/rancher/gitwatcher/pkg/generated/controllers/gitwatcher.cattle.io/v1"
	riov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	"github.com/rancher/rio/pkg/constants"
	riov1controller "github.com/rancher/rio/pkg/generated/controllers/rio.cattle.io/v1"
	"github.com/rancher/rio/types"
	appsv1controller "github.com/rancher/wrangler-api/pkg/generated/controllers/apps/v1"
)

type Handler struct {
	ctx              context.Context
	namespace        string
	gitWatcherCache  webhookv1controller.GitWatcherCache
	gitWatcherClient webhookv1controller.GitWatcherClient
	gitcommits       webhookv1controller.GitCommitController
	deploymentClient appsv1controller.DeploymentClient
	deploymentCache  appsv1controller.DeploymentCache
	dwClient         riov1controller.DeploymentWranglerController
	sswClient        riov1controller.StatefulSetWranglerController
	stacks           riov1controller.StackController
}

func Register(ctx context.Context, rContext *types.Context) error {
	h := Handler{
		ctx:              ctx,
		namespace:        rContext.Namespace,
		deploymentClient: rContext.Apps.Apps().V1().Deployment(),
		deploymentCache:  rContext.Apps.Apps().V1().Deployment().Cache(),
		dwClient:         rContext.Rio.Rio().V1().DeploymentWrangler(),
		sswClient:        rContext.Rio.Rio().V1().StatefulSetWrangler(),
		stacks:           rContext.Rio.Rio().V1().Stack(),
		gitWatcherCache:  rContext.Webhook.Gitwatcher().V1().GitWatcher().Cache(),
		gitWatcherClient: rContext.Webhook.Gitwatcher().V1().GitWatcher(),
		gitcommits:       rContext.Webhook.Gitwatcher().V1().GitCommit(),
	}

	rContext.Webhook.Gitwatcher().V1().GitCommit().OnChange(ctx, "webhook-execution", h.onChange)

	rContext.Rio.Rio().V1().DeploymentWrangler().OnChange(ctx, "service-update-gitcommit", h.dwUpdateGitCommit)
	rContext.Rio.Rio().V1().StatefulSetWrangler().OnChange(ctx, "service-update-gitcommit", h.sswUpdateGitCommit)

	return nil
}

// onChange gets fired when a new gitcommit comes in
func (h Handler) onChange(key string, gc *webhookv1.GitCommit) (*webhookv1.GitCommit, error) {
	if gc == nil {
		return gc, nil
	}

	if webhookv1.GitWebHookExecutionConditionHandled.IsTrue(gc) {
		return gc, nil
	}

	gitWatcher, err := h.gitWatcherCache.Get(gc.Namespace, gc.Spec.GitWatcherName)
	if err != nil {
		return nil, err
	}

	// todo: fix stack
	//if isOwnedByStack(gitWatcher) {
	//	if _, err := h.onGitcommitUpdateStack(key, obj, gitWatcher); err != nil {
	//		return nil, err
	//	}
	//}

	if _, err := h.onGitCommitChangeUpdateWorkload(key, gc, gitWatcher); err != nil {
		return nil, err
	}

	gc = gc.DeepCopy()
	webhookv1.GitWebHookExecutionConditionHandled.SetStatus(gc, "True")
	_, err = h.gitcommits.Update(gc)
	return gc, err
}

func (h Handler) dwUpdateGitCommit(key string, dw *riov1.DeploymentWrangler) (*riov1.DeploymentWrangler, error) {
	if dw == nil || dw.DeletionTimestamp != nil {
		return dw, nil
	}
	_, err := h.updateGitCommit(key, dw)
	return dw, err
}

func (h Handler) sswUpdateGitCommit(key string, ssw *riov1.StatefulSetWrangler) (*riov1.StatefulSetWrangler, error) {
	if ssw == nil || ssw.DeletionTimestamp != nil {
		return ssw, nil
	}
	_, err := h.updateGitCommit(key, ssw)
	return ssw, err
}

func (h Handler) updateGitCommit(key string, w riov1.Workload) (riov1.Workload, error) {

	if w.GetMeta().Annotations[constants.GitCommitLabel] == "" {
		return w, nil
	}
	//
	//rev := w.GetSpec().ImageBuild.Revisionw
	//if rev == "" {
	//	return w, nil
	//}
	//
	//gitcommit, err := h.gitcommits.Cache().Get(w.GetMeta().Namespace, w.GetMeta().Annotations[constants.GitCommitLabel])
	//if err != nil {
	//	return w, err
	//}
	//
	//if gitcommit.Status.GithubStatus == nil {
	//	return w, nil
	//}
	//
	//webhook, err := h.services.Cache().Get(h.namespace, "webhook")
	//if err != nil {
	//	return w, err
	//}
	//
	//webhookEndpoint := ""
	//if len(webhook.Status.Endpoints) > 0 {
	//	webhookEndpoint = webhook.Status.Endpoints[0]
	//}
	//
	//gitcommit = gitcommit.DeepCopy()
	//logURL := fmt.Sprintf("%s/logs/%s/%s?log-token=%s", webhookEndpoint, w.Namespace, w.Name, w.Status.BuildLogToken)
	//endpoint := ""
	//if len(w.Status.Endpoints) > 0 {
	//	endpoint = w.Status.Endpoints[0]
	//}
	//state := "in_progress"
	//if w.Status.DeploymentReady {
	//	state = "success"
	//}
	//update := false
	//if gitcommit.Status.GithubStatus.LogURL != logURL || gitcommit.Status.GithubStatus.EnvironmentURL != endpoint || gitcommit.Status.GithubStatus.DeploymentState != state {
	//	update = true
	//}
	//if !update {
	//	return w, nil
	//}
	//
	//gitcommit.Status.GithubStatus.LogURL = logURL
	//gitcommit.Status.GithubStatus.EnvironmentURL = endpoint
	//gitcommit.Status.GithubStatus.DeploymentState = state
	//if _, err := h.gitcommits.Update(gitcommit); err != nil {
	//	return w, err
	//}
	return w, nil
}

func isOwnedByStack(gitWatcher *webhookv1.GitWatcher) bool {
	return gitWatcher.Annotations["objectset.rio.cattle.io/owner-gvk"] == "rio.cattle.io/v1, Kind=Stack"
}
