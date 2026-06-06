package ctlstoragev1

import (
	wranglerstoragev1 "github.com/rancher/wrangler/v3/pkg/generated/controllers/storage/v1"
)

// Types not yet in Harvester generated controllers — shim from wrangler
type StorageClassController = wranglerstoragev1.StorageClassController
type StorageClassCache = wranglerstoragev1.StorageClassCache
type StorageClassClient = wranglerstoragev1.StorageClassClient
