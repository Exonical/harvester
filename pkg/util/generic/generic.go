package generic

import (
	wranglergeneric "github.com/rancher/wrangler/v3/pkg/generic"
	"k8s.io/apimachinery/pkg/runtime"
)

type RuntimeMetaObject = wranglergeneric.RuntimeMetaObject
type ClientInterface[T RuntimeMetaObject, TList runtime.Object] = wranglergeneric.ClientInterface[T, TList]
type ControllerInterface[T RuntimeMetaObject, TList runtime.Object] = wranglergeneric.ControllerInterface[T, TList]
type CacheInterface[T RuntimeMetaObject] = wranglergeneric.CacheInterface[T]
type NonNamespacedControllerInterface[T RuntimeMetaObject, TList runtime.Object] = wranglergeneric.NonNamespacedControllerInterface[T, TList]
type NonNamespacedCacheInterface[T RuntimeMetaObject] = wranglergeneric.NonNamespacedCacheInterface[T]
