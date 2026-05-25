package util

import corev1 "k8s.io/api/core/v1"

const (
	VipConfigmapName = "vip"
)

// VIPConfig represents the VIP configuration stored in a ConfigMap
type VIPConfig struct {
	Enabled        string             `json:"enabled,omitempty"`
	ServiceType    corev1.ServiceType `json:"serviceType,omitempty"`
	IP             string             `json:"ip,omitempty"`
	Mode           string             `json:"mode,omitempty"`
	HwAddress      string             `json:"hwAddress,omitempty"`
	LoadBalancerIP string             `json:"loadBalancerIP,omitempty"`
}
