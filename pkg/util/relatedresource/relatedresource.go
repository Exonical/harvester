package relatedresource

import (
	wranglerrelated "github.com/rancher/wrangler/v3/pkg/relatedresource"
)

type Key = wranglerrelated.Key
type Resolver = wranglerrelated.Resolver
type ControllerWrapper = wranglerrelated.ControllerWrapper
type Enqueuer = wranglerrelated.Enqueuer
type ClusterScopedEnqueuer = wranglerrelated.ClusterScopedEnqueuer

var Watch = wranglerrelated.Watch
var WatchClusterScoped = wranglerrelated.WatchClusterScoped
var NewKey = wranglerrelated.NewKey
