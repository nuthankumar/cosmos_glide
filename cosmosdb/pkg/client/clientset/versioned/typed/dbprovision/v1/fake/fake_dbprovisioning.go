/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	dbprovisionv1 "cosmosdb/pkg/apis/dbprovision/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDBProvisionings implements DBProvisioningInterface
type FakeDBProvisionings struct {
	Fake *FakeDbprovisionV1
	ns   string
}

var dbprovisioningsResource = schema.GroupVersionResource{Group: "dbprovision.com", Version: "v1", Resource: "dbprovisionings"}

var dbprovisioningsKind = schema.GroupVersionKind{Group: "dbprovision.com", Version: "v1", Kind: "DBProvisioning"}

// Get takes name of the dBProvisioning, and returns the corresponding dBProvisioning object, and an error if there is any.
func (c *FakeDBProvisionings) Get(ctx context.Context, name string, options v1.GetOptions) (result *dbprovisionv1.DBProvisioning, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dbprovisioningsResource, c.ns, name), &dbprovisionv1.DBProvisioning{})

	if obj == nil {
		return nil, err
	}
	return obj.(*dbprovisionv1.DBProvisioning), err
}

// List takes label and field selectors, and returns the list of DBProvisionings that match those selectors.
func (c *FakeDBProvisionings) List(ctx context.Context, opts v1.ListOptions) (result *dbprovisionv1.DBProvisioningList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dbprovisioningsResource, dbprovisioningsKind, c.ns, opts), &dbprovisionv1.DBProvisioningList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &dbprovisionv1.DBProvisioningList{ListMeta: obj.(*dbprovisionv1.DBProvisioningList).ListMeta}
	for _, item := range obj.(*dbprovisionv1.DBProvisioningList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dBProvisionings.
func (c *FakeDBProvisionings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dbprovisioningsResource, c.ns, opts))

}

// Create takes the representation of a dBProvisioning and creates it.  Returns the server's representation of the dBProvisioning, and an error, if there is any.
func (c *FakeDBProvisionings) Create(ctx context.Context, dBProvisioning *dbprovisionv1.DBProvisioning, opts v1.CreateOptions) (result *dbprovisionv1.DBProvisioning, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dbprovisioningsResource, c.ns, dBProvisioning), &dbprovisionv1.DBProvisioning{})

	if obj == nil {
		return nil, err
	}
	return obj.(*dbprovisionv1.DBProvisioning), err
}

// Update takes the representation of a dBProvisioning and updates it. Returns the server's representation of the dBProvisioning, and an error, if there is any.
func (c *FakeDBProvisionings) Update(ctx context.Context, dBProvisioning *dbprovisionv1.DBProvisioning, opts v1.UpdateOptions) (result *dbprovisionv1.DBProvisioning, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dbprovisioningsResource, c.ns, dBProvisioning), &dbprovisionv1.DBProvisioning{})

	if obj == nil {
		return nil, err
	}
	return obj.(*dbprovisionv1.DBProvisioning), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDBProvisionings) UpdateStatus(ctx context.Context, dBProvisioning *dbprovisionv1.DBProvisioning, opts v1.UpdateOptions) (*dbprovisionv1.DBProvisioning, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dbprovisioningsResource, "status", c.ns, dBProvisioning), &dbprovisionv1.DBProvisioning{})

	if obj == nil {
		return nil, err
	}
	return obj.(*dbprovisionv1.DBProvisioning), err
}

// Delete takes name of the dBProvisioning and deletes it. Returns an error if one occurs.
func (c *FakeDBProvisionings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dbprovisioningsResource, c.ns, name), &dbprovisionv1.DBProvisioning{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDBProvisionings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dbprovisioningsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &dbprovisionv1.DBProvisioningList{})
	return err
}

// Patch applies the patch and returns the patched dBProvisioning.
func (c *FakeDBProvisionings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *dbprovisionv1.DBProvisioning, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dbprovisioningsResource, c.ns, name, pt, data, subresources...), &dbprovisionv1.DBProvisioning{})

	if obj == nil {
		return nil, err
	}
	return obj.(*dbprovisionv1.DBProvisioning), err
}
