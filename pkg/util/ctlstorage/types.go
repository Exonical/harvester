package ctlstorage

import (
	wstorage "github.com/rancher/wrangler/v3/pkg/generated/controllers/storage"
	wstoragev1 "github.com/rancher/wrangler/v3/pkg/generated/controllers/storage/v1"
	"github.com/rancher/wrangler/v3/pkg/generic"
	"k8s.io/client-go/rest"
)

// Factory is a type alias for the wrangler storage controller factory.
type Factory = wstorage.Factory

// NewFactoryFromConfigWithOptions creates a new storage Factory.
func NewFactoryFromConfigWithOptions(config *rest.Config, opts *generic.FactoryOptions) (*Factory, error) {
	return wstorage.NewFactoryFromConfigWithOptions(config, opts)
}

// Controller type aliases for storage/v1 resources.
type (
	StorageClassCache  = wstoragev1.StorageClassCache
	StorageClassClient = wstoragev1.StorageClassClient
)
