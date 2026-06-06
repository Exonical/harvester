package apply

import (
	wapply "github.com/rancher/wrangler/v3/pkg/apply"
	"k8s.io/client-go/rest"
)

// Apply is a type alias for the wrangler Apply interface.
// This package centralizes the wrangler apply dependency so application code
// does not import wrangler directly.
type Apply = wapply.Apply

// NewForConfig creates a new Apply instance from a rest.Config.
func NewForConfig(cfg *rest.Config) (Apply, error) {
	return wapply.NewForConfig(cfg)
}
