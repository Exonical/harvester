// Package corev1 re-exports wrangler-generated core/v1 controller cache and
// client types so that consuming packages inside pkg/api/ and pkg/server/ no
// longer need a direct wrangler import.
package corev1

import wranglercorev1 "github.com/rancher/wrangler/v3/pkg/generated/controllers/core/v1"

type PodCache = wranglercorev1.PodCache
type SecretCache = wranglercorev1.SecretCache
type NodeCache = wranglercorev1.NodeCache
type NodeClient = wranglercorev1.NodeClient
type NamespaceCache = wranglercorev1.NamespaceCache
type ConfigMapCache = wranglercorev1.ConfigMapCache
type PersistentVolumeCache = wranglercorev1.PersistentVolumeCache
type PersistentVolumeClaimCache = wranglercorev1.PersistentVolumeClaimCache
type PersistentVolumeClaimClient = wranglercorev1.PersistentVolumeClaimClient
type SecretClient = wranglercorev1.SecretClient
type ServiceAccountClient = wranglercorev1.ServiceAccountClient
type ServiceAccountCache = wranglercorev1.ServiceAccountCache
