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
	gitsecretv1 "github.com/harshajith/k8s-controller-secrets-crd/pkg/apis/gitsecret/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGitSecrets implements GitSecretInterface
type FakeGitSecrets struct {
	Fake *FakeGitsecretV1
	ns   string
}

var gitsecretsResource = schema.GroupVersionResource{Group: "gitsecret.com", Version: "v1", Resource: "gitsecrets"}

var gitsecretsKind = schema.GroupVersionKind{Group: "gitsecret.com", Version: "v1", Kind: "GitSecret"}

// Get takes name of the gitSecret, and returns the corresponding gitSecret object, and an error if there is any.
func (c *FakeGitSecrets) Get(name string, options v1.GetOptions) (result *gitsecretv1.GitSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(gitsecretsResource, c.ns, name), &gitsecretv1.GitSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gitsecretv1.GitSecret), err
}

// List takes label and field selectors, and returns the list of GitSecrets that match those selectors.
func (c *FakeGitSecrets) List(opts v1.ListOptions) (result *gitsecretv1.GitSecretList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(gitsecretsResource, gitsecretsKind, c.ns, opts), &gitsecretv1.GitSecretList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &gitsecretv1.GitSecretList{ListMeta: obj.(*gitsecretv1.GitSecretList).ListMeta}
	for _, item := range obj.(*gitsecretv1.GitSecretList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gitSecrets.
func (c *FakeGitSecrets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(gitsecretsResource, c.ns, opts))

}

// Create takes the representation of a gitSecret and creates it.  Returns the server's representation of the gitSecret, and an error, if there is any.
func (c *FakeGitSecrets) Create(gitSecret *gitsecretv1.GitSecret) (result *gitsecretv1.GitSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(gitsecretsResource, c.ns, gitSecret), &gitsecretv1.GitSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gitsecretv1.GitSecret), err
}

// Update takes the representation of a gitSecret and updates it. Returns the server's representation of the gitSecret, and an error, if there is any.
func (c *FakeGitSecrets) Update(gitSecret *gitsecretv1.GitSecret) (result *gitsecretv1.GitSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(gitsecretsResource, c.ns, gitSecret), &gitsecretv1.GitSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gitsecretv1.GitSecret), err
}

// Delete takes name of the gitSecret and deletes it. Returns an error if one occurs.
func (c *FakeGitSecrets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(gitsecretsResource, c.ns, name), &gitsecretv1.GitSecret{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGitSecrets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(gitsecretsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &gitsecretv1.GitSecretList{})
	return err
}

// Patch applies the patch and returns the patched gitSecret.
func (c *FakeGitSecrets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *gitsecretv1.GitSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gitsecretsResource, c.ns, name, pt, data, subresources...), &gitsecretv1.GitSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gitsecretv1.GitSecret), err
}