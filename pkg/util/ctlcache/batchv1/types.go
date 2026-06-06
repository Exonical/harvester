// Package batchv1 re-exports wrangler-generated batch/v1 controller cache
// types so that consuming packages no longer need a direct wrangler import.
package batchv1

import wranglerbatchv1 "github.com/rancher/wrangler/v3/pkg/generated/controllers/batch/v1"

type JobCache = wranglerbatchv1.JobCache
type JobClient = wranglerbatchv1.JobClient
