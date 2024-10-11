/*
Copyright The KubeVirt Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha2 "kubevirt.io/api/instancetype/v1alpha2"
)

// FakeVirtualMachinePreferences implements VirtualMachinePreferenceInterface
type FakeVirtualMachinePreferences struct {
	Fake *FakeInstancetypeV1alpha2
	ns   string
}

var virtualmachinepreferencesResource = v1alpha2.SchemeGroupVersion.WithResource("virtualmachinepreferences")

var virtualmachinepreferencesKind = v1alpha2.SchemeGroupVersion.WithKind("VirtualMachinePreference")

// Get takes name of the virtualMachinePreference, and returns the corresponding virtualMachinePreference object, and an error if there is any.
func (c *FakeVirtualMachinePreferences) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.VirtualMachinePreference, err error) {
	emptyResult := &v1alpha2.VirtualMachinePreference{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(virtualmachinepreferencesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.VirtualMachinePreference), err
}

// List takes label and field selectors, and returns the list of VirtualMachinePreferences that match those selectors.
func (c *FakeVirtualMachinePreferences) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.VirtualMachinePreferenceList, err error) {
	emptyResult := &v1alpha2.VirtualMachinePreferenceList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(virtualmachinepreferencesResource, virtualmachinepreferencesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.VirtualMachinePreferenceList{ListMeta: obj.(*v1alpha2.VirtualMachinePreferenceList).ListMeta}
	for _, item := range obj.(*v1alpha2.VirtualMachinePreferenceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualMachinePreferences.
func (c *FakeVirtualMachinePreferences) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(virtualmachinepreferencesResource, c.ns, opts))

}

// Create takes the representation of a virtualMachinePreference and creates it.  Returns the server's representation of the virtualMachinePreference, and an error, if there is any.
func (c *FakeVirtualMachinePreferences) Create(ctx context.Context, virtualMachinePreference *v1alpha2.VirtualMachinePreference, opts v1.CreateOptions) (result *v1alpha2.VirtualMachinePreference, err error) {
	emptyResult := &v1alpha2.VirtualMachinePreference{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(virtualmachinepreferencesResource, c.ns, virtualMachinePreference, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.VirtualMachinePreference), err
}

// Update takes the representation of a virtualMachinePreference and updates it. Returns the server's representation of the virtualMachinePreference, and an error, if there is any.
func (c *FakeVirtualMachinePreferences) Update(ctx context.Context, virtualMachinePreference *v1alpha2.VirtualMachinePreference, opts v1.UpdateOptions) (result *v1alpha2.VirtualMachinePreference, err error) {
	emptyResult := &v1alpha2.VirtualMachinePreference{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(virtualmachinepreferencesResource, c.ns, virtualMachinePreference, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.VirtualMachinePreference), err
}

// Delete takes name of the virtualMachinePreference and deletes it. Returns an error if one occurs.
func (c *FakeVirtualMachinePreferences) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(virtualmachinepreferencesResource, c.ns, name, opts), &v1alpha2.VirtualMachinePreference{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualMachinePreferences) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(virtualmachinepreferencesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.VirtualMachinePreferenceList{})
	return err
}

// Patch applies the patch and returns the patched virtualMachinePreference.
func (c *FakeVirtualMachinePreferences) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.VirtualMachinePreference, err error) {
	emptyResult := &v1alpha2.VirtualMachinePreference{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(virtualmachinepreferencesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.VirtualMachinePreference), err
}