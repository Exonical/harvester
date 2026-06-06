package ctlbatch

import (
	wbatch "github.com/rancher/wrangler/v3/pkg/generated/controllers/batch"
	wbatchv1 "github.com/rancher/wrangler/v3/pkg/generated/controllers/batch/v1"
	"github.com/rancher/wrangler/v3/pkg/generic"
	"k8s.io/client-go/rest"
)

// Factory is a type alias for the wrangler batch controller factory.
type Factory = wbatch.Factory

// NewFactoryFromConfigWithOptions creates a new batch Factory.
func NewFactoryFromConfigWithOptions(config *rest.Config, opts *generic.FactoryOptions) (*Factory, error) {
	return wbatch.NewFactoryFromConfigWithOptions(config, opts)
}

// Controller type aliases for batch/v1 resources.
type (
	JobController = wbatchv1.JobController
)
