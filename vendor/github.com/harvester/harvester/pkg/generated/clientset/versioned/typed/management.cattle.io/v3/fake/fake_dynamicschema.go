/*
Copyright 2023 Rancher Labs, Inc.

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
	"context"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDynamicSchemas implements DynamicSchemaInterface
type FakeDynamicSchemas struct {
	Fake *FakeManagementV3
}

var dynamicschemasResource = schema.GroupVersionResource{Group: "management.cattle.io", Version: "v3", Resource: "dynamicschemas"}

var dynamicschemasKind = schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "DynamicSchema"}

// Get takes name of the dynamicSchema, and returns the corresponding dynamicSchema object, and an error if there is any.
func (c *FakeDynamicSchemas) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.DynamicSchema, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(dynamicschemasResource, name), &v3.DynamicSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.DynamicSchema), err
}

// List takes label and field selectors, and returns the list of DynamicSchemas that match those selectors.
func (c *FakeDynamicSchemas) List(ctx context.Context, opts v1.ListOptions) (result *v3.DynamicSchemaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(dynamicschemasResource, dynamicschemasKind, opts), &v3.DynamicSchemaList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.DynamicSchemaList{ListMeta: obj.(*v3.DynamicSchemaList).ListMeta}
	for _, item := range obj.(*v3.DynamicSchemaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dynamicSchemas.
func (c *FakeDynamicSchemas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(dynamicschemasResource, opts))
}

// Create takes the representation of a dynamicSchema and creates it.  Returns the server's representation of the dynamicSchema, and an error, if there is any.
func (c *FakeDynamicSchemas) Create(ctx context.Context, dynamicSchema *v3.DynamicSchema, opts v1.CreateOptions) (result *v3.DynamicSchema, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(dynamicschemasResource, dynamicSchema), &v3.DynamicSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.DynamicSchema), err
}

// Update takes the representation of a dynamicSchema and updates it. Returns the server's representation of the dynamicSchema, and an error, if there is any.
func (c *FakeDynamicSchemas) Update(ctx context.Context, dynamicSchema *v3.DynamicSchema, opts v1.UpdateOptions) (result *v3.DynamicSchema, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(dynamicschemasResource, dynamicSchema), &v3.DynamicSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.DynamicSchema), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDynamicSchemas) UpdateStatus(ctx context.Context, dynamicSchema *v3.DynamicSchema, opts v1.UpdateOptions) (*v3.DynamicSchema, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(dynamicschemasResource, "status", dynamicSchema), &v3.DynamicSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.DynamicSchema), err
}

// Delete takes name of the dynamicSchema and deletes it. Returns an error if one occurs.
func (c *FakeDynamicSchemas) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(dynamicschemasResource, name, opts), &v3.DynamicSchema{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDynamicSchemas) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(dynamicschemasResource, listOpts)

	_, err := c.Fake.Invokes(action, &v3.DynamicSchemaList{})
	return err
}

// Patch applies the patch and returns the patched dynamicSchema.
func (c *FakeDynamicSchemas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.DynamicSchema, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(dynamicschemasResource, name, pt, data, subresources...), &v3.DynamicSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.DynamicSchema), err
}