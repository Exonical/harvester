package gvk

import (
	"encoding/json"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Detect extracts the GVK from raw JSON bytes.
func Detect(obj []byte) (schema.GroupVersionKind, bool, error) {
	partial := v1.PartialObjectMetadata{}
	if err := json.Unmarshal(obj, &partial); err != nil {
		return schema.GroupVersionKind{}, false, err
	}
	result := partial.GetObjectKind().GroupVersionKind()
	ok := result.Kind != "" && result.Version != ""
	return result, ok, nil
}

// Get returns the GVK for a runtime.Object, looking it up in the scheme if needed.
func Get(scheme *runtime.Scheme, obj runtime.Object) (schema.GroupVersionKind, error) {
	gvk := obj.GetObjectKind().GroupVersionKind()
	if gvk.Kind != "" {
		return gvk, nil
	}

	gvks, _, err := scheme.ObjectKinds(obj)
	if err != nil {
		return schema.GroupVersionKind{}, fmt.Errorf("failed to find gvk for %T: %w", obj, err)
	}

	if len(gvks) == 0 {
		return schema.GroupVersionKind{}, fmt.Errorf("failed to find gvk for %T", obj)
	}

	return gvks[0], nil
}

// Set populates the GVK on runtime.Objects using the scheme.
func Set(scheme *runtime.Scheme, objs ...runtime.Object) error {
	for _, obj := range objs {
		if err := setObject(scheme, obj); err != nil {
			return err
		}
	}
	return nil
}

func setObject(scheme *runtime.Scheme, obj runtime.Object) error {
	gvk := obj.GetObjectKind().GroupVersionKind()
	if gvk.Kind != "" {
		return nil
	}

	gvks, _, err := scheme.ObjectKinds(obj)
	if err != nil {
		return err
	}

	if len(gvks) == 0 {
		return nil
	}

	obj.GetObjectKind().SetGroupVersionKind(gvks[0])
	return nil
}
