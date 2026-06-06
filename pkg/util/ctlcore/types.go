package ctlcore

import (
	wcore "github.com/rancher/wrangler/v3/pkg/generated/controllers/core"
	wcorev1 "github.com/rancher/wrangler/v3/pkg/generated/controllers/core/v1"
	"github.com/rancher/wrangler/v3/pkg/generic"
	"k8s.io/client-go/rest"
)

// Factory is a type alias for the wrangler core controller factory.
type Factory = wcore.Factory

// NewFactoryFromConfig creates a new core Factory with default options.
func NewFactoryFromConfig(config *rest.Config) (*Factory, error) {
	return wcore.NewFactoryFromConfig(config)
}

// NewFactoryFromConfigWithOptions creates a new core Factory.
func NewFactoryFromConfigWithOptions(config *rest.Config, opts *generic.FactoryOptions) (*Factory, error) {
	return wcore.NewFactoryFromConfigWithOptions(config, opts)
}

// Controller type aliases for core/v1 resources.
type (
	SecretClient                  = wcorev1.SecretClient
	SecretCache                   = wcorev1.SecretCache
	SecretController              = wcorev1.SecretController
	PersistentVolumeClaimCache    = wcorev1.PersistentVolumeClaimCache
	PersistentVolumeClaimClient   = wcorev1.PersistentVolumeClaimClient
	PersistentVolumeClaimController = wcorev1.PersistentVolumeClaimController
	PodCache                      = wcorev1.PodCache
	PodController                 = wcorev1.PodController
	ServiceController             = wcorev1.ServiceController
	NodeCache                     = wcorev1.NodeCache
	NodeClient                    = wcorev1.NodeClient
	EndpointsCache                = wcorev1.EndpointsCache
	NamespaceCache                = wcorev1.NamespaceCache
)
