/*
Copyright 2020 Rancher Labs.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	clientset "github.com/rancher/rio/pkg/generated/clientset/versioned/typed/rio.cattle.io/v1"
	informers "github.com/rancher/rio/pkg/generated/informers/externalversions/rio.cattle.io/v1"
	listers "github.com/rancher/rio/pkg/generated/listers/rio.cattle.io/v1"
	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type DeploymentWranglerHandler func(string, *v1.DeploymentWrangler) (*v1.DeploymentWrangler, error)

type DeploymentWranglerController interface {
	generic.ControllerMeta
	DeploymentWranglerClient

	OnChange(ctx context.Context, name string, sync DeploymentWranglerHandler)
	OnRemove(ctx context.Context, name string, sync DeploymentWranglerHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() DeploymentWranglerCache
}

type DeploymentWranglerClient interface {
	Create(*v1.DeploymentWrangler) (*v1.DeploymentWrangler, error)
	Update(*v1.DeploymentWrangler) (*v1.DeploymentWrangler, error)
	UpdateStatus(*v1.DeploymentWrangler) (*v1.DeploymentWrangler, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1.DeploymentWrangler, error)
	List(namespace string, opts metav1.ListOptions) (*v1.DeploymentWranglerList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.DeploymentWrangler, err error)
}

type DeploymentWranglerCache interface {
	Get(namespace, name string) (*v1.DeploymentWrangler, error)
	List(namespace string, selector labels.Selector) ([]*v1.DeploymentWrangler, error)

	AddIndexer(indexName string, indexer DeploymentWranglerIndexer)
	GetByIndex(indexName, key string) ([]*v1.DeploymentWrangler, error)
}

type DeploymentWranglerIndexer func(obj *v1.DeploymentWrangler) ([]string, error)

type deploymentWranglerController struct {
	controllerManager *generic.ControllerManager
	clientGetter      clientset.DeploymentWranglersGetter
	informer          informers.DeploymentWranglerInformer
	gvk               schema.GroupVersionKind
}

func NewDeploymentWranglerController(gvk schema.GroupVersionKind, controllerManager *generic.ControllerManager, clientGetter clientset.DeploymentWranglersGetter, informer informers.DeploymentWranglerInformer) DeploymentWranglerController {
	return &deploymentWranglerController{
		controllerManager: controllerManager,
		clientGetter:      clientGetter,
		informer:          informer,
		gvk:               gvk,
	}
}

func FromDeploymentWranglerHandlerToHandler(sync DeploymentWranglerHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.DeploymentWrangler
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.DeploymentWrangler))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *deploymentWranglerController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.DeploymentWrangler))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateDeploymentWranglerDeepCopyOnChange(client DeploymentWranglerClient, obj *v1.DeploymentWrangler, handler func(obj *v1.DeploymentWrangler) (*v1.DeploymentWrangler, error)) (*v1.DeploymentWrangler, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *deploymentWranglerController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controllerManager.AddHandler(ctx, c.gvk, c.informer.Informer(), name, handler)
}

func (c *deploymentWranglerController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	removeHandler := generic.NewRemoveHandler(name, c.Updater(), handler)
	c.controllerManager.AddHandler(ctx, c.gvk, c.informer.Informer(), name, removeHandler)
}

func (c *deploymentWranglerController) OnChange(ctx context.Context, name string, sync DeploymentWranglerHandler) {
	c.AddGenericHandler(ctx, name, FromDeploymentWranglerHandlerToHandler(sync))
}

func (c *deploymentWranglerController) OnRemove(ctx context.Context, name string, sync DeploymentWranglerHandler) {
	removeHandler := generic.NewRemoveHandler(name, c.Updater(), FromDeploymentWranglerHandlerToHandler(sync))
	c.AddGenericHandler(ctx, name, removeHandler)
}

func (c *deploymentWranglerController) Enqueue(namespace, name string) {
	c.controllerManager.Enqueue(c.gvk, c.informer.Informer(), namespace, name)
}

func (c *deploymentWranglerController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controllerManager.EnqueueAfter(c.gvk, c.informer.Informer(), namespace, name, duration)
}

func (c *deploymentWranglerController) Informer() cache.SharedIndexInformer {
	return c.informer.Informer()
}

func (c *deploymentWranglerController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *deploymentWranglerController) Cache() DeploymentWranglerCache {
	return &deploymentWranglerCache{
		lister:  c.informer.Lister(),
		indexer: c.informer.Informer().GetIndexer(),
	}
}

func (c *deploymentWranglerController) Create(obj *v1.DeploymentWrangler) (*v1.DeploymentWrangler, error) {
	return c.clientGetter.DeploymentWranglers(obj.Namespace).Create(obj)
}

func (c *deploymentWranglerController) Update(obj *v1.DeploymentWrangler) (*v1.DeploymentWrangler, error) {
	return c.clientGetter.DeploymentWranglers(obj.Namespace).Update(obj)
}

func (c *deploymentWranglerController) UpdateStatus(obj *v1.DeploymentWrangler) (*v1.DeploymentWrangler, error) {
	return c.clientGetter.DeploymentWranglers(obj.Namespace).UpdateStatus(obj)
}

func (c *deploymentWranglerController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return c.clientGetter.DeploymentWranglers(namespace).Delete(name, options)
}

func (c *deploymentWranglerController) Get(namespace, name string, options metav1.GetOptions) (*v1.DeploymentWrangler, error) {
	return c.clientGetter.DeploymentWranglers(namespace).Get(name, options)
}

func (c *deploymentWranglerController) List(namespace string, opts metav1.ListOptions) (*v1.DeploymentWranglerList, error) {
	return c.clientGetter.DeploymentWranglers(namespace).List(opts)
}

func (c *deploymentWranglerController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientGetter.DeploymentWranglers(namespace).Watch(opts)
}

func (c *deploymentWranglerController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.DeploymentWrangler, err error) {
	return c.clientGetter.DeploymentWranglers(namespace).Patch(name, pt, data, subresources...)
}

type deploymentWranglerCache struct {
	lister  listers.DeploymentWranglerLister
	indexer cache.Indexer
}

func (c *deploymentWranglerCache) Get(namespace, name string) (*v1.DeploymentWrangler, error) {
	return c.lister.DeploymentWranglers(namespace).Get(name)
}

func (c *deploymentWranglerCache) List(namespace string, selector labels.Selector) ([]*v1.DeploymentWrangler, error) {
	return c.lister.DeploymentWranglers(namespace).List(selector)
}

func (c *deploymentWranglerCache) AddIndexer(indexName string, indexer DeploymentWranglerIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.DeploymentWrangler))
		},
	}))
}

func (c *deploymentWranglerCache) GetByIndex(indexName, key string) (result []*v1.DeploymentWrangler, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		result = append(result, obj.(*v1.DeploymentWrangler))
	}
	return result, nil
}

type DeploymentWranglerStatusHandler func(obj *v1.DeploymentWrangler, status v1.DeploymentWranglerStatus) (v1.DeploymentWranglerStatus, error)

type DeploymentWranglerGeneratingHandler func(obj *v1.DeploymentWrangler, status v1.DeploymentWranglerStatus) ([]runtime.Object, v1.DeploymentWranglerStatus, error)

func RegisterDeploymentWranglerStatusHandler(ctx context.Context, controller DeploymentWranglerController, condition condition.Cond, name string, handler DeploymentWranglerStatusHandler) {
	statusHandler := &deploymentWranglerStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromDeploymentWranglerHandlerToHandler(statusHandler.sync))
}

func RegisterDeploymentWranglerGeneratingHandler(ctx context.Context, controller DeploymentWranglerController, apply apply.Apply,
	condition condition.Cond, name string, handler DeploymentWranglerGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &deploymentWranglerGeneratingHandler{
		DeploymentWranglerGeneratingHandler: handler,
		apply:                               apply,
		name:                                name,
		gvk:                                 controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	RegisterDeploymentWranglerStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type deploymentWranglerStatusHandler struct {
	client    DeploymentWranglerClient
	condition condition.Cond
	handler   DeploymentWranglerStatusHandler
}

func (a *deploymentWranglerStatusHandler) sync(key string, obj *v1.DeploymentWrangler) (*v1.DeploymentWrangler, error) {
	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {
		// Revert to old status on error
		newStatus = *origStatus.DeepCopy()
	}

	obj.Status = newStatus
	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(obj, "", nil)
		} else {
			a.condition.SetError(obj, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, obj.Status) {
		var newErr error
		obj, newErr = a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
	}
	return obj, err
}

type deploymentWranglerGeneratingHandler struct {
	DeploymentWranglerGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *deploymentWranglerGeneratingHandler) Handle(obj *v1.DeploymentWrangler, status v1.DeploymentWranglerStatus) (v1.DeploymentWranglerStatus, error) {
	objs, newStatus, err := a.DeploymentWranglerGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	apply := a.apply

	if !a.opts.DynamicLookup {
		apply = apply.WithStrictCaching()
	}

	if !a.opts.AllowCrossNamespace && !a.opts.AllowClusterScoped {
		apply = apply.WithSetOwnerReference(true, false).
			WithDefaultNamespace(obj.GetNamespace()).
			WithListerNamespace(obj.GetNamespace())
	}

	if !a.opts.AllowClusterScoped {
		apply = apply.WithRestrictClusterScoped()
	}

	if a.opts.WithoutOwnerReference {
		apply = apply.WithoutOwnerReference()
	}

	return newStatus, apply.
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
