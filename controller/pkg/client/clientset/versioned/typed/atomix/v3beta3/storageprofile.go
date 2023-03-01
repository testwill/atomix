// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by client-gen. DO NOT EDIT.

package v3beta3

import (
	"context"
	"time"

	v3beta3 "github.com/atomix/atomix/controller/pkg/apis/atomix/v3beta3"
	scheme "github.com/atomix/atomix/controller/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// StorageProfilesGetter has a method to return a StorageProfileInterface.
// A group's client should implement this interface.
type StorageProfilesGetter interface {
	StorageProfiles(namespace string) StorageProfileInterface
}

// StorageProfileInterface has methods to work with StorageProfile resources.
type StorageProfileInterface interface {
	Create(ctx context.Context, storageProfile *v3beta3.StorageProfile, opts v1.CreateOptions) (*v3beta3.StorageProfile, error)
	Update(ctx context.Context, storageProfile *v3beta3.StorageProfile, opts v1.UpdateOptions) (*v3beta3.StorageProfile, error)
	UpdateStatus(ctx context.Context, storageProfile *v3beta3.StorageProfile, opts v1.UpdateOptions) (*v3beta3.StorageProfile, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3beta3.StorageProfile, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3beta3.StorageProfileList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3beta3.StorageProfile, err error)
	StorageProfileExpansion
}

// storageProfiles implements StorageProfileInterface
type storageProfiles struct {
	client rest.Interface
	ns     string
}

// newStorageProfiles returns a StorageProfiles
func newStorageProfiles(c *AtomixV3beta3Client, namespace string) *storageProfiles {
	return &storageProfiles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the storageProfile, and returns the corresponding storageProfile object, and an error if there is any.
func (c *storageProfiles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3beta3.StorageProfile, err error) {
	result = &v3beta3.StorageProfile{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("storageprofiles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StorageProfiles that match those selectors.
func (c *storageProfiles) List(ctx context.Context, opts v1.ListOptions) (result *v3beta3.StorageProfileList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3beta3.StorageProfileList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("storageprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested storageProfiles.
func (c *storageProfiles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("storageprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a storageProfile and creates it.  Returns the server's representation of the storageProfile, and an error, if there is any.
func (c *storageProfiles) Create(ctx context.Context, storageProfile *v3beta3.StorageProfile, opts v1.CreateOptions) (result *v3beta3.StorageProfile, err error) {
	result = &v3beta3.StorageProfile{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("storageprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(storageProfile).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a storageProfile and updates it. Returns the server's representation of the storageProfile, and an error, if there is any.
func (c *storageProfiles) Update(ctx context.Context, storageProfile *v3beta3.StorageProfile, opts v1.UpdateOptions) (result *v3beta3.StorageProfile, err error) {
	result = &v3beta3.StorageProfile{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("storageprofiles").
		Name(storageProfile.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(storageProfile).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *storageProfiles) UpdateStatus(ctx context.Context, storageProfile *v3beta3.StorageProfile, opts v1.UpdateOptions) (result *v3beta3.StorageProfile, err error) {
	result = &v3beta3.StorageProfile{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("storageprofiles").
		Name(storageProfile.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(storageProfile).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the storageProfile and deletes it. Returns an error if one occurs.
func (c *storageProfiles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("storageprofiles").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *storageProfiles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("storageprofiles").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched storageProfile.
func (c *storageProfiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3beta3.StorageProfile, err error) {
	result = &v3beta3.StorageProfile{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("storageprofiles").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
