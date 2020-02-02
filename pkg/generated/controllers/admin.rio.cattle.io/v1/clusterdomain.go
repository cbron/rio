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

	v1 "github.com/rancher/rio/pkg/apis/admin.rio.cattle.io/v1"
	clientset "github.com/rancher/rio/pkg/generated/clientset/versioned/typed/admin.rio.cattle.io/v1"
	informers "github.com/rancher/rio/pkg/generated/informers/externalversions/admin.rio.cattle.io/v1"
	listers "github.com/rancher/rio/pkg/generated/listers/admin.rio.cattle.io/v1"
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

type ClusterDomainHandler func(string, *v1.ClusterDomain) (*v1.ClusterDomain, error)

type ClusterDomainController interface {
	generic.ControllerMeta
	ClusterDomainClient

	OnChange(ctx context.Context, name string, sync ClusterDomainHandler)
	OnRemove(ctx context.Context, name string, sync ClusterDomainHandler)
	Enqueue(name string)
	EnqueueAfter(name string, duration time.Duration)

	Cache() ClusterDomainCache
}

type ClusterDomainClient interface {
	Create(*v1.ClusterDomain) (*v1.ClusterDomain, error)
	Update(*v1.ClusterDomain) (*v1.ClusterDomain, error)
	UpdateStatus(*v1.ClusterDomain) (*v1.ClusterDomain, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ClusterDomain, error)
	List(opts metav1.ListOptions) (*v1.ClusterDomainList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClusterDomain, err error)
}

type ClusterDomainCache interface {
	Get(name string) (*v1.ClusterDomain, error)
	List(selector labels.Selector) ([]*v1.ClusterDomain, error)

	AddIndexer(indexName string, indexer ClusterDomainIndexer)
	GetByIndex(indexName, key string) ([]*v1.ClusterDomain, error)
}

type ClusterDomainIndexer func(obj *v1.ClusterDomain) ([]string, error)

type clusterDomainController struct {
	controllerManager *generic.ControllerManager
	clientGetter      clientset.ClusterDomainsGetter
	informer          informers.ClusterDomainInformer
	gvk               schema.GroupVersionKind
}

func NewClusterDomainController(gvk schema.GroupVersionKind, controllerManager *generic.ControllerManager, clientGetter clientset.ClusterDomainsGetter, informer informers.ClusterDomainInformer) ClusterDomainController {
	return &clusterDomainController{
		controllerManager: controllerManager,
		clientGetter:      clientGetter,
		informer:          informer,
		gvk:               gvk,
	}
}

func FromClusterDomainHandlerToHandler(sync ClusterDomainHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.ClusterDomain
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.ClusterDomain))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *clusterDomainController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.ClusterDomain))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateClusterDomainDeepCopyOnChange(client ClusterDomainClient, obj *v1.ClusterDomain, handler func(obj *v1.ClusterDomain) (*v1.ClusterDomain, error)) (*v1.ClusterDomain, error) {
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

func (c *clusterDomainController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controllerManager.AddHandler(ctx, c.gvk, c.informer.Informer(), name, handler)
}

func (c *clusterDomainController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	removeHandler := generic.NewRemoveHandler(name, c.Updater(), handler)
	c.controllerManager.AddHandler(ctx, c.gvk, c.informer.Informer(), name, removeHandler)
}

func (c *clusterDomainController) OnChange(ctx context.Context, name string, sync ClusterDomainHandler) {
	c.AddGenericHandler(ctx, name, FromClusterDomainHandlerToHandler(sync))
}

func (c *clusterDomainController) OnRemove(ctx context.Context, name string, sync ClusterDomainHandler) {
	removeHandler := generic.NewRemoveHandler(name, c.Updater(), FromClusterDomainHandlerToHandler(sync))
	c.AddGenericHandler(ctx, name, removeHandler)
}

func (c *clusterDomainController) Enqueue(name string) {
	c.controllerManager.Enqueue(c.gvk, c.informer.Informer(), "", name)
}

func (c *clusterDomainController) EnqueueAfter(name string, duration time.Duration) {
	c.controllerManager.EnqueueAfter(c.gvk, c.informer.Informer(), "", name, duration)
}

func (c *clusterDomainController) Informer() cache.SharedIndexInformer {
	return c.informer.Informer()
}

func (c *clusterDomainController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *clusterDomainController) Cache() ClusterDomainCache {
	return &clusterDomainCache{
		lister:  c.informer.Lister(),
		indexer: c.informer.Informer().GetIndexer(),
	}
}

func (c *clusterDomainController) Create(obj *v1.ClusterDomain) (*v1.ClusterDomain, error) {
	return c.clientGetter.ClusterDomains().Create(obj)
}

func (c *clusterDomainController) Update(obj *v1.ClusterDomain) (*v1.ClusterDomain, error) {
	return c.clientGetter.ClusterDomains().Update(obj)
}

func (c *clusterDomainController) UpdateStatus(obj *v1.ClusterDomain) (*v1.ClusterDomain, error) {
	return c.clientGetter.ClusterDomains().UpdateStatus(obj)
}

func (c *clusterDomainController) Delete(name string, options *metav1.DeleteOptions) error {
	return c.clientGetter.ClusterDomains().Delete(name, options)
}

func (c *clusterDomainController) Get(name string, options metav1.GetOptions) (*v1.ClusterDomain, error) {
	return c.clientGetter.ClusterDomains().Get(name, options)
}

func (c *clusterDomainController) List(opts metav1.ListOptions) (*v1.ClusterDomainList, error) {
	return c.clientGetter.ClusterDomains().List(opts)
}

func (c *clusterDomainController) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientGetter.ClusterDomains().Watch(opts)
}

func (c *clusterDomainController) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClusterDomain, err error) {
	return c.clientGetter.ClusterDomains().Patch(name, pt, data, subresources...)
}

type clusterDomainCache struct {
	lister  listers.ClusterDomainLister
	indexer cache.Indexer
}

func (c *clusterDomainCache) Get(name string) (*v1.ClusterDomain, error) {
	return c.lister.Get(name)
}

func (c *clusterDomainCache) List(selector labels.Selector) ([]*v1.ClusterDomain, error) {
	return c.lister.List(selector)
}

func (c *clusterDomainCache) AddIndexer(indexName string, indexer ClusterDomainIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.ClusterDomain))
		},
	}))
}

func (c *clusterDomainCache) GetByIndex(indexName, key string) (result []*v1.ClusterDomain, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		result = append(result, obj.(*v1.ClusterDomain))
	}
	return result, nil
}

type ClusterDomainStatusHandler func(obj *v1.ClusterDomain, status v1.ClusterDomainStatus) (v1.ClusterDomainStatus, error)

type ClusterDomainGeneratingHandler func(obj *v1.ClusterDomain, status v1.ClusterDomainStatus) ([]runtime.Object, v1.ClusterDomainStatus, error)

func RegisterClusterDomainStatusHandler(ctx context.Context, controller ClusterDomainController, condition condition.Cond, name string, handler ClusterDomainStatusHandler) {
	statusHandler := &clusterDomainStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromClusterDomainHandlerToHandler(statusHandler.sync))
}

func RegisterClusterDomainGeneratingHandler(ctx context.Context, controller ClusterDomainController, apply apply.Apply,
	condition condition.Cond, name string, handler ClusterDomainGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &clusterDomainGeneratingHandler{
		ClusterDomainGeneratingHandler: handler,
		apply:                          apply,
		name:                           name,
		gvk:                            controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	RegisterClusterDomainStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type clusterDomainStatusHandler struct {
	client    ClusterDomainClient
	condition condition.Cond
	handler   ClusterDomainStatusHandler
}

func (a *clusterDomainStatusHandler) sync(key string, obj *v1.ClusterDomain) (*v1.ClusterDomain, error) {
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

type clusterDomainGeneratingHandler struct {
	ClusterDomainGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *clusterDomainGeneratingHandler) Handle(obj *v1.ClusterDomain, status v1.ClusterDomainStatus) (v1.ClusterDomainStatus, error) {
	objs, newStatus, err := a.ClusterDomainGeneratingHandler(obj, status)
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
