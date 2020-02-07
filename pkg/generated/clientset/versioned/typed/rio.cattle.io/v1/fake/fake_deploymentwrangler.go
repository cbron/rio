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

package fake

import (
	riocattleiov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDeploymentWranglers implements DeploymentWranglerInterface
type FakeDeploymentWranglers struct {
	Fake *FakeRioV1
	ns   string
}

var deploymentwranglersResource = schema.GroupVersionResource{Group: "rio.cattle.io", Version: "v1", Resource: "deploymentwranglers"}

var deploymentwranglersKind = schema.GroupVersionKind{Group: "rio.cattle.io", Version: "v1", Kind: "DeploymentWrangler"}

// Get takes name of the deploymentWrangler, and returns the corresponding deploymentWrangler object, and an error if there is any.
func (c *FakeDeploymentWranglers) Get(name string, options v1.GetOptions) (result *riocattleiov1.DeploymentWrangler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(deploymentwranglersResource, c.ns, name), &riocattleiov1.DeploymentWrangler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*riocattleiov1.DeploymentWrangler), err
}

// List takes label and field selectors, and returns the list of DeploymentWranglers that match those selectors.
func (c *FakeDeploymentWranglers) List(opts v1.ListOptions) (result *riocattleiov1.DeploymentWranglerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(deploymentwranglersResource, deploymentwranglersKind, c.ns, opts), &riocattleiov1.DeploymentWranglerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &riocattleiov1.DeploymentWranglerList{ListMeta: obj.(*riocattleiov1.DeploymentWranglerList).ListMeta}
	for _, item := range obj.(*riocattleiov1.DeploymentWranglerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested deploymentWranglers.
func (c *FakeDeploymentWranglers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(deploymentwranglersResource, c.ns, opts))

}

// Create takes the representation of a deploymentWrangler and creates it.  Returns the server's representation of the deploymentWrangler, and an error, if there is any.
func (c *FakeDeploymentWranglers) Create(deploymentWrangler *riocattleiov1.DeploymentWrangler) (result *riocattleiov1.DeploymentWrangler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(deploymentwranglersResource, c.ns, deploymentWrangler), &riocattleiov1.DeploymentWrangler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*riocattleiov1.DeploymentWrangler), err
}

// Update takes the representation of a deploymentWrangler and updates it. Returns the server's representation of the deploymentWrangler, and an error, if there is any.
func (c *FakeDeploymentWranglers) Update(deploymentWrangler *riocattleiov1.DeploymentWrangler) (result *riocattleiov1.DeploymentWrangler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(deploymentwranglersResource, c.ns, deploymentWrangler), &riocattleiov1.DeploymentWrangler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*riocattleiov1.DeploymentWrangler), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDeploymentWranglers) UpdateStatus(deploymentWrangler *riocattleiov1.DeploymentWrangler) (*riocattleiov1.DeploymentWrangler, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(deploymentwranglersResource, "status", c.ns, deploymentWrangler), &riocattleiov1.DeploymentWrangler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*riocattleiov1.DeploymentWrangler), err
}

// Delete takes name of the deploymentWrangler and deletes it. Returns an error if one occurs.
func (c *FakeDeploymentWranglers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(deploymentwranglersResource, c.ns, name), &riocattleiov1.DeploymentWrangler{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDeploymentWranglers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(deploymentwranglersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &riocattleiov1.DeploymentWranglerList{})
	return err
}

// Patch applies the patch and returns the patched deploymentWrangler.
func (c *FakeDeploymentWranglers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *riocattleiov1.DeploymentWrangler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(deploymentwranglersResource, c.ns, name, pt, data, subresources...), &riocattleiov1.DeploymentWrangler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*riocattleiov1.DeploymentWrangler), err
}