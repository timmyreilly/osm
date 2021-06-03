/*
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

package v1alpha1

import (
	v1alpha1 "github.com/openservicemesh/osm/pkg/apis/config/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RemoteServiceLister helps list RemoteServices.
// All objects returned here must be treated as read-only.
type RemoteServiceLister interface {
	// List lists all RemoteServices in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RemoteService, err error)
	// RemoteServices returns an object that can list and get RemoteServices.
	RemoteServices(namespace string) RemoteServiceNamespaceLister
	RemoteServiceListerExpansion
}

// remoteServiceLister implements the RemoteServiceLister interface.
type remoteServiceLister struct {
	indexer cache.Indexer
}

// NewRemoteServiceLister returns a new RemoteServiceLister.
func NewRemoteServiceLister(indexer cache.Indexer) RemoteServiceLister {
	return &remoteServiceLister{indexer: indexer}
}

// List lists all RemoteServices in the indexer.
func (s *remoteServiceLister) List(selector labels.Selector) (ret []*v1alpha1.RemoteService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RemoteService))
	})
	return ret, err
}

// RemoteServices returns an object that can list and get RemoteServices.
func (s *remoteServiceLister) RemoteServices(namespace string) RemoteServiceNamespaceLister {
	return remoteServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RemoteServiceNamespaceLister helps list and get RemoteServices.
// All objects returned here must be treated as read-only.
type RemoteServiceNamespaceLister interface {
	// List lists all RemoteServices in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RemoteService, err error)
	// Get retrieves the RemoteService from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.RemoteService, error)
	RemoteServiceNamespaceListerExpansion
}

// remoteServiceNamespaceLister implements the RemoteServiceNamespaceLister
// interface.
type remoteServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RemoteServices in the indexer for a given namespace.
func (s remoteServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RemoteService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RemoteService))
	})
	return ret, err
}

// Get retrieves the RemoteService from the indexer for a given namespace and name.
func (s remoteServiceNamespaceLister) Get(name string) (*v1alpha1.RemoteService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("remoteservice"), name)
	}
	return obj.(*v1alpha1.RemoteService), nil
}
