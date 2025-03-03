/*
Copyright 2021 The Knative Authors

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
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
	v1 "knative.dev/eventing/pkg/apis/flows/v1"
)

// ParallelLister helps list Parallels.
// All objects returned here must be treated as read-only.
type ParallelLister interface {
	// List lists all Parallels in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Parallel, err error)
	// Parallels returns an object that can list and get Parallels.
	Parallels(namespace string) ParallelNamespaceLister
	ParallelListerExpansion
}

// parallelLister implements the ParallelLister interface.
type parallelLister struct {
	listers.ResourceIndexer[*v1.Parallel]
}

// NewParallelLister returns a new ParallelLister.
func NewParallelLister(indexer cache.Indexer) ParallelLister {
	return &parallelLister{listers.New[*v1.Parallel](indexer, v1.Resource("parallel"))}
}

// Parallels returns an object that can list and get Parallels.
func (s *parallelLister) Parallels(namespace string) ParallelNamespaceLister {
	return parallelNamespaceLister{listers.NewNamespaced[*v1.Parallel](s.ResourceIndexer, namespace)}
}

// ParallelNamespaceLister helps list and get Parallels.
// All objects returned here must be treated as read-only.
type ParallelNamespaceLister interface {
	// List lists all Parallels in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Parallel, err error)
	// Get retrieves the Parallel from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Parallel, error)
	ParallelNamespaceListerExpansion
}

// parallelNamespaceLister implements the ParallelNamespaceLister
// interface.
type parallelNamespaceLister struct {
	listers.ResourceIndexer[*v1.Parallel]
}
