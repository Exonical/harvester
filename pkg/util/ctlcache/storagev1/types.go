// Package storagev1 re-exports wrangler-generated storage/v1 controller cache
// types so that consuming packages no longer need a direct wrangler import.
package storagev1

import wranglerstoragev1 "github.com/rancher/wrangler/v3/pkg/generated/controllers/storage/v1"

type StorageClassCache = wranglerstoragev1.StorageClassCache
