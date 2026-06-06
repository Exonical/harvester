package ctlcorev1

import (
	harvcorev1 "github.com/harvester/harvester/pkg/generated/controllers/core/v1"
	wranglercorev1 "github.com/rancher/wrangler/v3/pkg/generated/controllers/core/v1"
)

// Types available in Harvester generated controllers
type NodeController = harvcorev1.NodeController
type NodeCache = harvcorev1.NodeCache
type NodeClient = harvcorev1.NodeClient
type PersistentVolumeController = harvcorev1.PersistentVolumeController
type PersistentVolumeCache = harvcorev1.PersistentVolumeCache
type PersistentVolumeClient = harvcorev1.PersistentVolumeClient
type PersistentVolumeClaimController = harvcorev1.PersistentVolumeClaimController
type PersistentVolumeClaimCache = harvcorev1.PersistentVolumeClaimCache
type PersistentVolumeClaimClient = harvcorev1.PersistentVolumeClaimClient
type ResourceQuotaController = harvcorev1.ResourceQuotaController
type ResourceQuotaCache = harvcorev1.ResourceQuotaCache
type ResourceQuotaClient = harvcorev1.ResourceQuotaClient

// Types not yet in Harvester generated controllers — shim from wrangler
type SecretController = wranglercorev1.SecretController
type SecretCache = wranglercorev1.SecretCache
type SecretClient = wranglercorev1.SecretClient
type PodCache = wranglercorev1.PodCache
type PodClient = wranglercorev1.PodClient
type ConfigMapCache = wranglercorev1.ConfigMapCache
type ConfigMapClient = wranglercorev1.ConfigMapClient
type ServiceClient = wranglercorev1.ServiceClient
type NamespaceController = wranglercorev1.NamespaceController
type NamespaceCache = wranglercorev1.NamespaceCache
type NamespaceClient = wranglercorev1.NamespaceClient
type EndpointsCache = wranglercorev1.EndpointsCache
