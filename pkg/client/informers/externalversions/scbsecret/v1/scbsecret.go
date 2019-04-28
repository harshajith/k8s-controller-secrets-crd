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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	scbsecretv1 "github.com/harshajith/k8s-controller-secrets-crd/pkg/apis/scbsecret/v1"
	versioned "github.com/harshajith/k8s-controller-secrets-crd/pkg/client/clientset/versioned"
	internalinterfaces "github.com/harshajith/k8s-controller-secrets-crd/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/harshajith/k8s-controller-secrets-crd/pkg/client/listers/scbsecret/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ScbSecretInformer provides access to a shared informer and lister for
// ScbSecrets.
type ScbSecretInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ScbSecretLister
}

type scbSecretInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewScbSecretInformer constructs a new informer for ScbSecret type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewScbSecretInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredScbSecretInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredScbSecretInformer constructs a new informer for ScbSecret type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredScbSecretInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ScbV1().ScbSecrets(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ScbV1().ScbSecrets(namespace).Watch(options)
			},
		},
		&scbsecretv1.ScbSecret{},
		resyncPeriod,
		indexers,
	)
}

func (f *scbSecretInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredScbSecretInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *scbSecretInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&scbsecretv1.ScbSecret{}, f.defaultInformer)
}

func (f *scbSecretInformer) Lister() v1.ScbSecretLister {
	return v1.NewScbSecretLister(f.Informer().GetIndexer())
}
