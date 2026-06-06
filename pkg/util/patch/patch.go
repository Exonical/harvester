package patch

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/evanphx/json-patch"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"k8s.io/client-go/kubernetes/scheme"

	"github.com/harvester/harvester/pkg/util/gvk"
)

var (
	patchCache     = map[schema.GroupVersionKind]patchCacheEntry{}
	patchCacheLock = sync.Mutex{}
)

type patchCacheEntry struct {
	patchType types.PatchType
	lookup    strategicpatch.LookupPatchMeta
}

func isJSONPatch(patch []byte) bool {
	return len(patch) > 0 && patch[0] == '['
}

// GetPatchStyle detects the appropriate patch type for the given original and patch.
func GetPatchStyle(original, patchData []byte) (types.PatchType, strategicpatch.LookupPatchMeta, error) {
	if isJSONPatch(patchData) {
		return types.JSONPatchType, nil, nil
	}
	gvkResult, ok, err := gvk.Detect(original)
	if err != nil {
		return "", nil, err
	}
	if !ok {
		return types.MergePatchType, nil, nil
	}
	return GetMergeStyle(gvkResult)
}

// GetMergeStyle returns the merge patch type and lookup metadata for a GVK.
func GetMergeStyle(gvkVal schema.GroupVersionKind) (types.PatchType, strategicpatch.LookupPatchMeta, error) {
	var (
		patchType       types.PatchType
		lookupPatchMeta strategicpatch.LookupPatchMeta
	)

	patchCacheLock.Lock()
	entry, ok := patchCache[gvkVal]
	patchCacheLock.Unlock()

	if ok {
		return entry.patchType, entry.lookup, nil
	}

	versionedObject, err := scheme.Scheme.New(gvkVal)

	if runtime.IsNotRegisteredError(err) || gvkVal.Kind == "CustomResourceDefinition" {
		patchType = types.MergePatchType
	} else if err != nil {
		return patchType, nil, err
	} else {
		patchType = types.StrategicMergePatchType
		lookupPatchMeta, err = strategicpatch.NewPatchMetaFromStruct(versionedObject)
		if err != nil {
			return patchType, nil, err
		}
	}

	patchCacheLock.Lock()
	patchCache[gvkVal] = patchCacheEntry{
		patchType: patchType,
		lookup:    lookupPatchMeta,
	}
	patchCacheLock.Unlock()

	return patchType, lookupPatchMeta, nil
}

// Apply applies a patch to the original JSON data.
func Apply(original, patchData []byte) ([]byte, error) {
	style, metadata, err := GetPatchStyle(original, patchData)
	if err != nil {
		return nil, err
	}

	switch style {
	case types.JSONPatchType:
		return applyJSONPatch(original, patchData)
	case types.MergePatchType:
		return applyMergePatch(original, patchData)
	case types.StrategicMergePatchType:
		return applyStrategicMergePatch(original, patchData, metadata)
	default:
		return nil, fmt.Errorf("invalid patch")
	}
}

func applyStrategicMergePatch(original, patchData []byte, lookup strategicpatch.LookupPatchMeta) ([]byte, error) {
	originalMap := map[string]interface{}{}
	patchMap := map[string]interface{}{}
	if err := json.Unmarshal(original, &originalMap); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(patchData, &patchMap); err != nil {
		return nil, err
	}
	patchedMap, err := strategicpatch.StrategicMergeMapPatchUsingLookupPatchMeta(originalMap, patchMap, lookup)
	if err != nil {
		return nil, err
	}
	return json.Marshal(patchedMap)
}

func applyMergePatch(original, patchData []byte) ([]byte, error) {
	return jsonpatch.MergePatch(original, patchData)
}

func applyJSONPatch(original, patchData []byte) ([]byte, error) {
	jp, err := jsonpatch.DecodePatch(patchData)
	if err != nil {
		return nil, err
	}
	return jp.Apply(original)
}
