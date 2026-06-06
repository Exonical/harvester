package ctlappsv1

import (
	harvappsv1 "github.com/harvester/harvester/pkg/generated/controllers/apps/v1"
	wranglerappsv1 "github.com/rancher/wrangler/v3/pkg/generated/controllers/apps/v1"
)

// Types available in Harvester generated controllers
type DeploymentController = harvappsv1.DeploymentController
type DeploymentCache = harvappsv1.DeploymentCache
type DeploymentClient = harvappsv1.DeploymentClient

// Types not yet in Harvester generated controllers — shim from wrangler
type DaemonSetCache = wranglerappsv1.DaemonSetCache
type DaemonSetClient = wranglerappsv1.DaemonSetClient
type StatefulSetCache = wranglerappsv1.StatefulSetCache
type StatefulSetClient = wranglerappsv1.StatefulSetClient
