// Package crt provides controller-runtime based client wrappers that replace
// wrangler's generated controller interfaces with a unified client approach.
package crt

import (
	"context"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TypedClient provides typed CRUD operations for a specific resource type,
// replacing wrangler's generated Client and Cache interfaces.
type TypedClient[T client.Object, L client.ObjectList] struct {
	client    client.Client
	namespace string
}

// NewTypedClient creates a TypedClient for the given resource type.
func NewTypedClient[T client.Object, L client.ObjectList](c client.Client, namespace string) *TypedClient[T, L] {
	return &TypedClient[T, L]{client: c, namespace: namespace}
}

func (tc *TypedClient[T, L]) Get(namespace, name string) (T, error) {
	var obj T
	obj = obj.DeepCopyObject().(T)
	err := tc.client.Get(context.Background(), types.NamespacedName{Namespace: namespace, Name: name}, obj)
	return obj, err
}

func (tc *TypedClient[T, L]) Create(obj T) (T, error) {
	err := tc.client.Create(context.Background(), obj)
	return obj, err
}

func (tc *TypedClient[T, L]) Update(obj T) (T, error) {
	err := tc.client.Update(context.Background(), obj)
	return obj, err
}

func (tc *TypedClient[T, L]) UpdateStatus(obj T) (T, error) {
	err := tc.client.Status().Update(context.Background(), obj)
	return obj, err
}

func (tc *TypedClient[T, L]) Delete(namespace, name string) error {
	var obj T
	obj = obj.DeepCopyObject().(T)
	obj.SetName(name)
	obj.SetNamespace(namespace)
	return tc.client.Delete(context.Background(), obj)
}

func (tc *TypedClient[T, L]) Patch(namespace, name string, pt types.PatchType, data []byte) (T, error) {
	var obj T
	obj = obj.DeepCopyObject().(T)
	err := tc.client.Get(context.Background(), types.NamespacedName{Namespace: namespace, Name: name}, obj)
	if err != nil {
		return obj, err
	}
	err = tc.client.Patch(context.Background(), obj, client.RawPatch(pt, data))
	return obj, err
}

func (tc *TypedClient[T, L]) List(namespace string, selector labels.Selector) (L, error) {
	var list L
	list = list.DeepCopyObject().(L)
	opts := &client.ListOptions{Namespace: namespace}
	if selector != nil {
		opts.LabelSelector = selector
	}
	err := tc.client.List(context.Background(), list, opts)
	return list, err
}
