// Package schemas re-exports wrangler schema types so that consuming
// packages no longer need a direct wrangler import.
package schemas

import wranglerschemas "github.com/rancher/wrangler/v3/pkg/schemas"

type Action = wranglerschemas.Action
