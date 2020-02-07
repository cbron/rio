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
	"time"

	v1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	scheme "github.com/rancher/rio/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DeploymentWranglersGetter has a method to return a DeploymentWranglerInterface.
// A group's client should implement this interface.
type DeploymentWranglersGetter interface {
	DeploymentWranglers(namespace string) DeploymentWranglerInterface
}

// DeploymentWranglerInterface has methods to work with DeploymentWrangler resources.
type DeploymentWranglerInterface interface {
	Create(*v1.DeploymentWrangler) (*v1.DeploymentWrangler, error)
	Update(*v1.DeploymentWrangler) (*v1.DeploymentWrangler, error)
	UpdateStatus(*v1.DeploymentWrangler) (*v1.DeploymentWrangler, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.DeploymentWrangler, error)
	List(opts metav1.ListOptions) (*v1.DeploymentWranglerList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.DeploymentWrangler, err error)
	DeploymentWranglerExpansion
}

// deploymentWranglers implements DeploymentWranglerInterface
type deploymentWranglers struct {
	client rest.Interface
	ns     string
}

// newDeploymentWranglers returns a DeploymentWranglers
func newDeploymentWranglers(c *RioV1Client, namespace string) *deploymentWranglers {
	return &deploymentWranglers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the deploymentWrangler, and returns the corresponding deploymentWrangler object, and an error if there is any.
func (c *deploymentWranglers) Get(name string, options metav1.GetOptions) (result *v1.DeploymentWrangler, err error) {
	result = &v1.DeploymentWrangler{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deploymentwranglers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DeploymentWranglers that match those selectors.
func (c *deploymentWranglers) List(opts metav1.ListOptions) (result *v1.DeploymentWranglerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DeploymentWranglerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deploymentwranglers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested deploymentWranglers.
func (c *deploymentWranglers) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("deploymentwranglers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a deploymentWrangler and creates it.  Returns the server's representation of the deploymentWrangler, and an error, if there is any.
func (c *deploymentWranglers) Create(deploymentWrangler *v1.DeploymentWrangler) (result *v1.DeploymentWrangler, err error) {
	result = &v1.DeploymentWrangler{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("deploymentwranglers").
		Body(deploymentWrangler).
		Do().
		Into(result)
	return
}

// Update takes the representation of a deploymentWrangler and updates it. Returns the server's representation of the deploymentWrangler, and an error, if there is any.
func (c *deploymentWranglers) Update(deploymentWrangler *v1.DeploymentWrangler) (result *v1.DeploymentWrangler, err error) {
	result = &v1.DeploymentWrangler{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deploymentwranglers").
		Name(deploymentWrangler.Name).
		Body(deploymentWrangler).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *deploymentWranglers) UpdateStatus(deploymentWrangler *v1.DeploymentWrangler) (result *v1.DeploymentWrangler, err error) {
	result = &v1.DeploymentWrangler{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deploymentwranglers").
		Name(deploymentWrangler.Name).
		SubResource("status").
		Body(deploymentWrangler).
		Do().
		Into(result)
	return
}

// Delete takes name of the deploymentWrangler and deletes it. Returns an error if one occurs.
func (c *deploymentWranglers) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deploymentwranglers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *deploymentWranglers) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deploymentwranglers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched deploymentWrangler.
func (c *deploymentWranglers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.DeploymentWrangler, err error) {
	result = &v1.DeploymentWrangler{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("deploymentwranglers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}