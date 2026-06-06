package generic

import (
	wgeneric "github.com/rancher/wrangler/v3/pkg/generic"
)

// FactoryOptions is a type alias for wrangler FactoryOptions.
// This package centralizes the wrangler generic dependency so application code
// does not import wrangler directly.
type FactoryOptions = wgeneric.FactoryOptions
