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
	v1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DeploymentWranglerLister helps list DeploymentWranglers.
type DeploymentWranglerLister interface {
	// List lists all DeploymentWranglers in the indexer.
	List(selector labels.Selector) (ret []*v1.DeploymentWrangler, err error)
	// DeploymentWranglers returns an object that can list and get DeploymentWranglers.
	DeploymentWranglers(namespace string) DeploymentWranglerNamespaceLister
	DeploymentWranglerListerExpansion
}

// deploymentWranglerLister implements the DeploymentWranglerLister interface.
type deploymentWranglerLister struct {
	indexer cache.Indexer
}

// NewDeploymentWranglerLister returns a new DeploymentWranglerLister.
func NewDeploymentWranglerLister(indexer cache.Indexer) DeploymentWranglerLister {
	return &deploymentWranglerLister{indexer: indexer}
}

// List lists all DeploymentWranglers in the indexer.
func (s *deploymentWranglerLister) List(selector labels.Selector) (ret []*v1.DeploymentWrangler, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DeploymentWrangler))
	})
	return ret, err
}

// DeploymentWranglers returns an object that can list and get DeploymentWranglers.
func (s *deploymentWranglerLister) DeploymentWranglers(namespace string) DeploymentWranglerNamespaceLister {
	return deploymentWranglerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DeploymentWranglerNamespaceLister helps list and get DeploymentWranglers.
type DeploymentWranglerNamespaceLister interface {
	// List lists all DeploymentWranglers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.DeploymentWrangler, err error)
	// Get retrieves the DeploymentWrangler from the indexer for a given namespace and name.
	Get(name string) (*v1.DeploymentWrangler, error)
	DeploymentWranglerNamespaceListerExpansion
}

// deploymentWranglerNamespaceLister implements the DeploymentWranglerNamespaceLister
// interface.
type deploymentWranglerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DeploymentWranglers in the indexer for a given namespace.
func (s deploymentWranglerNamespaceLister) List(selector labels.Selector) (ret []*v1.DeploymentWrangler, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DeploymentWrangler))
	})
	return ret, err
}

// Get retrieves the DeploymentWrangler from the indexer for a given namespace and name.
func (s deploymentWranglerNamespaceLister) Get(name string) (*v1.DeploymentWrangler, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("deploymentwrangler"), name)
	}
	return obj.(*v1.DeploymentWrangler), nil
}