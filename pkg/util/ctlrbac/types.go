package ctlrbac

import (
	wrbac "github.com/rancher/wrangler/v3/pkg/generated/controllers/rbac"
	"github.com/rancher/wrangler/v3/pkg/generic"
	"k8s.io/client-go/rest"
)

// Factory is a type alias for the wrangler rbac controller factory.
type Factory = wrbac.Factory

// NewFactoryFromConfigWithOptions creates a new rbac Factory.
func NewFactoryFromConfigWithOptions(config *rest.Config, opts *generic.FactoryOptions) (*Factory, error) {
	return wrbac.NewFactoryFromConfigWithOptions(config, opts)
}
