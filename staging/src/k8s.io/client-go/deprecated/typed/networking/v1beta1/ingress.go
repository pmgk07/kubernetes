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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "k8s.io/api/networking/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/deprecated/scheme"
	rest "k8s.io/client-go/rest"
)

// IngressesGetter has a method to return a IngressInterface.
// A group's client should implement this interface.
type IngressesGetter interface {
	Ingresses(namespace string) IngressInterface
}

// IngressInterface has methods to work with Ingress resources.
type IngressInterface interface {
	Create(*v1beta1.Ingress) (*v1beta1.Ingress, error)
	Update(*v1beta1.Ingress) (*v1beta1.Ingress, error)
	UpdateStatus(*v1beta1.Ingress) (*v1beta1.Ingress, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.Ingress, error)
	List(opts v1.ListOptions) (*v1beta1.IngressList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Ingress, err error)
	IngressExpansion
}

// ingresses implements IngressInterface
type ingresses struct {
	client rest.Interface
	ns     string
}

// newIngresses returns a Ingresses
func newIngresses(c *NetworkingV1beta1Client, namespace string) *ingresses {
	return &ingresses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the ingress, and returns the corresponding ingress object, and an error if there is any.
func (c *ingresses) Get(name string, options v1.GetOptions) (result *v1beta1.Ingress, err error) {
	result = &v1beta1.Ingress{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ingresses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Ingresses that match those selectors.
func (c *ingresses) List(opts v1.ListOptions) (result *v1beta1.IngressList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.IngressList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ingresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.TODO()).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ingresses.
func (c *ingresses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ingresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(context.TODO())
}

// Create takes the representation of a ingress and creates it.  Returns the server's representation of the ingress, and an error, if there is any.
func (c *ingresses) Create(ingress *v1beta1.Ingress) (result *v1beta1.Ingress, err error) {
	result = &v1beta1.Ingress{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ingresses").
		Body(ingress).
		Do(context.TODO()).
		Into(result)
	return
}

// Update takes the representation of a ingress and updates it. Returns the server's representation of the ingress, and an error, if there is any.
func (c *ingresses) Update(ingress *v1beta1.Ingress) (result *v1beta1.Ingress, err error) {
	result = &v1beta1.Ingress{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ingresses").
		Name(ingress.Name).
		Body(ingress).
		Do(context.TODO()).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *ingresses) UpdateStatus(ingress *v1beta1.Ingress) (result *v1beta1.Ingress, err error) {
	result = &v1beta1.Ingress{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ingresses").
		Name(ingress.Name).
		SubResource("status").
		Body(ingress).
		Do(context.TODO()).
		Into(result)
	return
}

// Delete takes name of the ingress and deletes it. Returns an error if one occurs.
func (c *ingresses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ingresses").
		Name(name).
		Body(options).
		Do(context.TODO()).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ingresses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ingresses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do(context.TODO()).
		Error()
}

// Patch applies the patch and returns the patched ingress.
func (c *ingresses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Ingress, err error) {
	result = &v1beta1.Ingress{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ingresses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.TODO()).
		Into(result)
	return
}