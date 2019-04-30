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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/harshajith/k8s-controller-secrets-crd/pkg/apis/gitsecret/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GitSecretLister helps list GitSecrets.
type GitSecretLister interface {
	// List lists all GitSecrets in the indexer.
	List(selector labels.Selector) (ret []*v1.GitSecret, err error)
	// GitSecrets returns an object that can list and get GitSecrets.
	GitSecrets(namespace string) GitSecretNamespaceLister
	GitSecretListerExpansion
}

// gitSecretLister implements the GitSecretLister interface.
type gitSecretLister struct {
	indexer cache.Indexer
}

// NewGitSecretLister returns a new GitSecretLister.
func NewGitSecretLister(indexer cache.Indexer) GitSecretLister {
	return &gitSecretLister{indexer: indexer}
}

// List lists all GitSecrets in the indexer.
func (s *gitSecretLister) List(selector labels.Selector) (ret []*v1.GitSecret, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.GitSecret))
	})
	return ret, err
}

// GitSecrets returns an object that can list and get GitSecrets.
func (s *gitSecretLister) GitSecrets(namespace string) GitSecretNamespaceLister {
	return gitSecretNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GitSecretNamespaceLister helps list and get GitSecrets.
type GitSecretNamespaceLister interface {
	// List lists all GitSecrets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.GitSecret, err error)
	// Get retrieves the GitSecret from the indexer for a given namespace and name.
	Get(name string) (*v1.GitSecret, error)
	GitSecretNamespaceListerExpansion
}

// gitSecretNamespaceLister implements the GitSecretNamespaceLister
// interface.
type gitSecretNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GitSecrets in the indexer for a given namespace.
func (s gitSecretNamespaceLister) List(selector labels.Selector) (ret []*v1.GitSecret, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.GitSecret))
	})
	return ret, err
}

// Get retrieves the GitSecret from the indexer for a given namespace and name.
func (s gitSecretNamespaceLister) Get(name string) (*v1.GitSecret, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("gitsecret"), name)
	}
	return obj.(*v1.GitSecret), nil
}