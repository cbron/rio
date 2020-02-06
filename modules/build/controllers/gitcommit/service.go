package gitcommit

import (
	//"fmt"

	"fmt"
	webhookv1 "github.com/rancher/gitwatcher/pkg/apis/gitwatcher.cattle.io/v1"
	riov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	"github.com/rancher/rio/pkg/constants"
	//"github.com/rancher/rio/pkg/indexes"
	"github.com/rancher/wrangler/pkg/name"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/api/errors"
)

func (h Handler) onGitCommitChangeUpdateWorkload(key string, gc *webhookv1.GitCommit, gitWatcher *webhookv1.GitWatcher) (*webhookv1.GitCommit, error) {
	if gc.Spec.Commit == "" && gc.Spec.Tag == "" {
		return gc, nil
	}
	ref := gc.Spec.Commit
	if ref == "" {
		ref = gc.Spec.Tag
	}

	if gitWatcher.Annotations[constants.WorkloadType] == constants.DeploymentWranglerType {
		return h.onGitCommitChangeUpdateDW(key, gc, ref, gitWatcher)
	} else if gitWatcher.Annotations[constants.WorkloadType] == constants.DeploymentWranglerType {
		// todo: implement ssw
	}
	return gc, nil
}

func (h Handler) onGitCommitChangeUpdateDW(key string, gc *webhookv1.GitCommit, ref string, gitWatcher *webhookv1.GitWatcher) (*webhookv1.GitCommit, error) {
	baseDW, err := h.dwClient.Cache().Get(gc.Namespace, gitWatcher.Annotations[constants.WorkloadName])
	if err != nil {
		if errors.IsNotFound(err) {
			return gc, nil
		}
		return gc, err
	}
	containerName := gitWatcher.Annotations[constants.ContainerLabel]
	var imageBuild *riov1.ImageBuildSpec
	index := 0
	for i, container := range baseDW.Spec.Containers {
		if container.Name == containerName {
			imageBuild = container.ImageBuild.DeepCopy()
			index = i
			break
		}
	}
	if imageBuild == nil {
		return gc, fmt.Errorf("container with name %s not found  in %s", containerName, baseDW.Name)
	}
	baseDW = baseDW.DeepCopy()

	// if its a template service, or an incoming PR commit
	if baseDW.Spec.Template || (gitWatcher.Spec.PR && gc.Spec.PR != "") {
		// if git commit is from different branch do no-op
		if gc.Spec.Branch != "" && gc.Spec.Branch != imageBuild.Branch {
			return gc, nil
		}
		baseDW.Status.WorkloadStatus = updateWorkloadStatus(baseDW, gc, ref, containerName, baseDW.Namespace)
		if _, err := h.dwClient.UpdateStatus(baseDW); err != nil {
			return gc, err
		}
	} else {
		// if template is false, just update existing images
		if imageBuild.Revision != ref {
			imageBuild.Revision = ref
			baseDW.Spec.Containers[index].ImageBuild = imageBuild
			if _, err := h.dwClient.Update(baseDW); err != nil {
				return gc, err
			}
		}
	}

	return gc, nil
}

func workloadName(w riov1.Workload, obj *webhookv1.GitCommit, ref string) string {
	n := name.SafeConcatName(w.GetMeta().Name, name.Hex(obj.Spec.RepositoryURL, 7))
	if obj.Spec.PR != "" {
		n = name.SafeConcatName(n, "pr"+obj.Spec.PR)
	} else {
		n = name.SafeConcatName(n, name.Hex(ref, 5))
	}
	return n
}

func updateWorkloadStatus(w riov1.Workload, gc *webhookv1.GitCommit, ref, containerName, ns string) riov1.WorkloadStatus {
	newName := workloadName(w, gc, ref)
	status := w.GetStatus()
	if status.ContainerRevision == nil {
		status.ContainerRevision = map[string]riov1.BuildRevision{}
	}
	status.GitCommits = append(status.GitCommits, gc.Name)

	if gc.Spec.PR != "" && (gc.Spec.Merged || gc.Spec.Closed) {
		logrus.Infof("PR %s is merged/closed, deleting revision, name: %s, namespace: %s, revision: %s", gc.Spec.PR, newName, ns, ref)
		if status.ShouldClean == nil {
			status.ShouldClean = map[string]bool{}
		}
		status.ShouldClean[newName] = true
	} else {
		revision := status.ContainerRevision[containerName]
		revision.Commits = append(revision.Commits, ref)
		status.ContainerRevision[containerName] = revision
		status.ShouldGenerate = newName
	}
	return status
}

//
//func (h Handler) scaleDownRevisions(namespace, name string) error {
//	revisions, err := h.services.Cache().GetByIndex(indexes.ServiceByApp, fmt.Sprintf("%s/%s", namespace, name))
//	if err != nil {
//		return err
//	}
//	for _, revision := range revisions {
//		deepcopy := revision.DeepCopy()
//		deepcopy.Spec.Weight = &[]int{0}[0]
//		if _, err := h.services.Update(deepcopy); err != nil {
//			return err
//		}
//		logrus.Infof("Scaling down service %s weight to 0", revision.Name)
//	}
//	return nil
//}
