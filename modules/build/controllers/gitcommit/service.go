package gitcommit

//
//func (h Handler) onChangeService(key string, obj *webhookv1.GitCommit, gitWatcher *webhookv1.GitWatcher) (*webhookv1.GitCommit, error) {
//	if obj.Spec.Commit == "" && obj.Spec.Tag == "" {
//		return obj, nil
//	}
//	ref := obj.Spec.Commit
//	if ref == "" {
//		ref = obj.Spec.Tag
//	}
//
//	baseService, err := h.services.Cache().Get(obj.Namespace, gitWatcher.Annotations[constants.ServiceLabel])
//	if err != nil {
//		if errors.IsNotFound(err) {
//			return obj, nil
//		}
//		return obj, err
//	}
//
//	var imageBuild riov1.ImageBuildSpec
//	containers := append(baseService.Spec.Containers, riov1.NamedContainer{
//		Name:      services.RootContainerName(baseService),
//		Container: baseService.Spec.Container,
//	})
//	containerName := gitWatcher.Annotations[constants.ContainerLabel]
//	for _, container := range containers {
//		if container.Name == containerName {
//			imageBuild = *container.ImageBuild
//			break
//		}
//	}
//
//	baseService = baseService.DeepCopy()
//
//	// if its a template service, or an incoming PR commit
//	if baseService.Spec.Template || (gitWatcher.Spec.PR && obj.Spec.PR != "") {
//		// if git commit is from different branch do no-op
//		if obj.Spec.Branch != "" && obj.Spec.Branch != imageBuild.Branch {
//			return obj, nil
//		}
//
//		serviceName := serviceName(baseService, obj, ref)
//		if baseService.Status.ContainerRevision == nil {
//			baseService.Status.ContainerRevision = map[string]riov1.BuildRevision{}
//		}
//		baseService.Status.GitCommits = append(baseService.Status.GitCommits, obj.Name)
//
//		if obj.Spec.PR != "" && (obj.Spec.Merged || obj.Spec.Closed) {
//			logrus.Infof("PR %s is merged/closed, deleting revision, name: %s, namespace: %s, revision: %s", obj.Spec.PR, serviceName, baseService.Namespace, ref)
//			if baseService.Status.ShouldClean == nil {
//				baseService.Status.ShouldClean = map[string]bool{}
//			}
//			baseService.Status.ShouldClean[serviceName] = true
//		} else {
//			revision := baseService.Status.ContainerRevision[containerName]
//			revision.Commits = append(revision.Commits, ref)
//			baseService.Status.ContainerRevision[containerName] = revision
//			baseService.Status.ShouldGenerate = serviceName
//		}
//
//		if _, err := h.services.UpdateStatus(baseService); err != nil {
//			return obj, err
//		}
//	} else {
//		// if template is false, just update existing images
//		update := false
//		if containerName == gitWatcher.Annotations[constants.ServiceLabel] {
//			// use the first container
//			if obj.Spec.Branch != "" && obj.Spec.Branch != baseService.Spec.ImageBuild.Branch {
//				return obj, nil
//			}
//			if baseService.Spec.ImageBuild.Revision != ref {
//				baseService.Spec.ImageBuild.Revision = ref
//				update = true
//			}
//		} else {
//			for i, con := range baseService.Spec.Containers {
//				if con.Name == containerName {
//					if baseService.Spec.Containers[i].ImageBuild.Revision != ref {
//						baseService.Spec.Containers[i].ImageBuild.Revision = ref
//						update = true
//					}
//				}
//			}
//		}
//		if update {
//			if _, err := h.services.Update(baseService); err != nil {
//				return obj, err
//			}
//		}
//	}
//
//	return obj, nil
//}
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
//
//func serviceName(service *riov1.Service, obj *webhookv1.GitCommit, ref string) string {
//	n := name.SafeConcatName(service.Name, name.Hex(obj.Spec.RepositoryURL, 7))
//	if obj.Spec.PR != "" {
//		n = name.SafeConcatName(n, "pr"+obj.Spec.PR)
//	} else {
//		n = name.SafeConcatName(n, name.Hex(ref, 5))
//	}
//	return n
//}
