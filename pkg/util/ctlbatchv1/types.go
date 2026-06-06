package ctlbatchv1

import (
	wranglerbatchv1 "github.com/rancher/wrangler/v3/pkg/generated/controllers/batch/v1"
)

// Types not yet in Harvester generated controllers — shim from wrangler
type JobController = wranglerbatchv1.JobController
type JobCache = wranglerbatchv1.JobCache
type JobClient = wranglerbatchv1.JobClient
