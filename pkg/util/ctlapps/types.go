package ctlapps

import (
	wapps "github.com/rancher/wrangler/v3/pkg/generated/controllers/apps"
	wappsv1 "github.com/rancher/wrangler/v3/pkg/generated/controllers/apps/v1"
	"github.com/rancher/wrangler/v3/pkg/generic"
	"k8s.io/client-go/rest"
)

// Factory is a type alias for the wrangler apps controller factory.
type Factory = wapps.Factory

// NewFactoryFromConfig creates a new apps Factory with default options.
func NewFactoryFromConfig(config *rest.Config) (*Factory, error) {
	return wapps.NewFactoryFromConfig(config)
}

// NewFactoryFromConfigWithOptions creates a new apps Factory.
func NewFactoryFromConfigWithOptions(config *rest.Config, opts *generic.FactoryOptions) (*Factory, error) {
	return wapps.NewFactoryFromConfigWithOptions(config, opts)
}

// Controller type aliases for apps/v1 resources.
type (
	DaemonSetCache = wappsv1.DaemonSetCache
)
