// Package generic re-exports wrangler generic types so that consuming packages
// no longer need a direct wrangler import.
package generic

import wranglergeneric "github.com/rancher/wrangler/v3/pkg/generic"

type FactoryOptions = wranglergeneric.FactoryOptions
