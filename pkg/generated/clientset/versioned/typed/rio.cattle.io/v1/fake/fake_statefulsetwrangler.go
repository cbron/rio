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

// FakeStatefulSetWranglers implements StatefulSetWranglerInterface
type FakeStatefulSetWranglers struct {
	Fake *FakeRioV1
	ns   string
}

var statefulsetwranglersResource = schema.GroupVersionResource{Group: "rio.cattle.io", Version: "v1", Resource: "statefulsetwranglers"}

var statefulsetwranglersKind = schema.GroupVersionKind{Group: "rio.cattle.io", Version: "v1", Kind: "StatefulSetWrangler"}

// Get takes name of the statefulSetWrangler, and returns the corresponding statefulSetWrangler object, and an error if there is any.
func (c *FakeStatefulSetWranglers) Get(name string, options v1.GetOptions) (result *riocattleiov1.StatefulSetWrangler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(statefulsetwranglersResource, c.ns, name), &riocattleiov1.StatefulSetWrangler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*riocattleiov1.StatefulSetWrangler), err
}

// List takes label and field selectors, and returns the list of StatefulSetWranglers that match those selectors.
func (c *FakeStatefulSetWranglers) List(opts v1.ListOptions) (result *riocattleiov1.StatefulSetWranglerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(statefulsetwranglersResource, statefulsetwranglersKind, c.ns, opts), &riocattleiov1.StatefulSetWranglerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &riocattleiov1.StatefulSetWranglerList{ListMeta: obj.(*riocattleiov1.StatefulSetWranglerList).ListMeta}
	for _, item := range obj.(*riocattleiov1.StatefulSetWranglerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested statefulSetWranglers.
func (c *FakeStatefulSetWranglers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(statefulsetwranglersResource, c.ns, opts))

}

// Create takes the representation of a statefulSetWrangler and creates it.  Returns the server's representation of the statefulSetWrangler, and an error, if there is any.
func (c *FakeStatefulSetWranglers) Create(statefulSetWrangler *riocattleiov1.StatefulSetWrangler) (result *riocattleiov1.StatefulSetWrangler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(statefulsetwranglersResource, c.ns, statefulSetWrangler), &riocattleiov1.StatefulSetWrangler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*riocattleiov1.StatefulSetWrangler), err
}

// Update takes the representation of a statefulSetWrangler and updates it. Returns the server's representation of the statefulSetWrangler, and an error, if there is any.
func (c *FakeStatefulSetWranglers) Update(statefulSetWrangler *riocattleiov1.StatefulSetWrangler) (result *riocattleiov1.StatefulSetWrangler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(statefulsetwranglersResource, c.ns, statefulSetWrangler), &riocattleiov1.StatefulSetWrangler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*riocattleiov1.StatefulSetWrangler), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStatefulSetWranglers) UpdateStatus(statefulSetWrangler *riocattleiov1.StatefulSetWrangler) (*riocattleiov1.StatefulSetWrangler, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(statefulsetwranglersResource, "status", c.ns, statefulSetWrangler), &riocattleiov1.StatefulSetWrangler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*riocattleiov1.StatefulSetWrangler), err
}

// Delete takes name of the statefulSetWrangler and deletes it. Returns an error if one occurs.
func (c *FakeStatefulSetWranglers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(statefulsetwranglersResource, c.ns, name), &riocattleiov1.StatefulSetWrangler{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStatefulSetWranglers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(statefulsetwranglersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &riocattleiov1.StatefulSetWranglerList{})
	return err
}

// Patch applies the patch and returns the patched statefulSetWrangler.
func (c *FakeStatefulSetWranglers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *riocattleiov1.StatefulSetWrangler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(statefulsetwranglersResource, c.ns, name, pt, data, subresources...), &riocattleiov1.StatefulSetWrangler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*riocattleiov1.StatefulSetWrangler), err
}
