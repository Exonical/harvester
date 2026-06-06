// Package generic provides local type aliases for the wrangler generic
// controller interfaces used by fakeclients. It isolates the wrangler
// dependency to a single file so that individual fakeclient files no
// longer import wrangler directly. Once the generated controller code
// is also migrated, these aliases can be replaced with standalone
// interface definitions.
package generic

import (
	wranglergeneric "github.com/rancher/wrangler/v3/pkg/generic"
	"k8s.io/apimachinery/pkg/runtime"
)

// RuntimeMetaObject mirrors the wrangler constraint for K8s objects.
type RuntimeMetaObject = wranglergeneric.RuntimeMetaObject

// ClientInterface performs CRUD operations on namespaced K8s objects.
type ClientInterface[T RuntimeMetaObject, TList runtime.Object] = wranglergeneric.ClientInterface[T, TList]

// NonNamespacedClientInterface performs CRUD operations on cluster-scoped K8s objects.
type NonNamespacedClientInterface[T RuntimeMetaObject, TList runtime.Object] = wranglergeneric.NonNamespacedClientInterface[T, TList]

// CacheInterface provides namespaced object retrieval from an in-memory cache.
type CacheInterface[T runtime.Object] = wranglergeneric.CacheInterface[T]

// NonNamespacedCacheInterface provides cluster-scoped object retrieval from an in-memory cache.
type NonNamespacedCacheInterface[T runtime.Object] = wranglergeneric.NonNamespacedCacheInterface[T]

// Indexer computes a set of indexed values for the provided object.
type Indexer[T runtime.Object] = wranglergeneric.Indexer[T]

// Handler processes a generic runtime.Object by key.
type Handler = wranglergeneric.Handler

// ObjectHandler processes a typed object by key.
type ObjectHandler[T runtime.Object] = wranglergeneric.ObjectHandler[T]

// Updater applies an update to a runtime.Object.
type Updater = wranglergeneric.Updater
