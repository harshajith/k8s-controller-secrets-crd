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
	v1 "github.com/harshajith/k8s-controller-secrets-crd/pkg/apis/scbsecret/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ScbSecretLister helps list ScbSecrets.
type ScbSecretLister interface {
	// List lists all ScbSecrets in the indexer.
	List(selector labels.Selector) (ret []*v1.ScbSecret, err error)
	// ScbSecrets returns an object that can list and get ScbSecrets.
	ScbSecrets(namespace string) ScbSecretNamespaceLister
	ScbSecretListerExpansion
}

// scbSecretLister implements the ScbSecretLister interface.
type scbSecretLister struct {
	indexer cache.Indexer
}

// NewScbSecretLister returns a new ScbSecretLister.
func NewScbSecretLister(indexer cache.Indexer) ScbSecretLister {
	return &scbSecretLister{indexer: indexer}
}

// List lists all ScbSecrets in the indexer.
func (s *scbSecretLister) List(selector labels.Selector) (ret []*v1.ScbSecret, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ScbSecret))
	})
	return ret, err
}

// ScbSecrets returns an object that can list and get ScbSecrets.
func (s *scbSecretLister) ScbSecrets(namespace string) ScbSecretNamespaceLister {
	return scbSecretNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ScbSecretNamespaceLister helps list and get ScbSecrets.
type ScbSecretNamespaceLister interface {
	// List lists all ScbSecrets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.ScbSecret, err error)
	// Get retrieves the ScbSecret from the indexer for a given namespace and name.
	Get(name string) (*v1.ScbSecret, error)
	ScbSecretNamespaceListerExpansion
}

// scbSecretNamespaceLister implements the ScbSecretNamespaceLister
// interface.
type scbSecretNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ScbSecrets in the indexer for a given namespace.
func (s scbSecretNamespaceLister) List(selector labels.Selector) (ret []*v1.ScbSecret, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ScbSecret))
	})
	return ret, err
}

// Get retrieves the ScbSecret from the indexer for a given namespace and name.
func (s scbSecretNamespaceLister) Get(name string) (*v1.ScbSecret, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("scbsecret"), name)
	}
	return obj.(*v1.ScbSecret), nil
}