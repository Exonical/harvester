package server

import (
	"fmt"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetConfig(kubeConfig string) (clientcmd.ClientConfig, error) {
	if isManual(kubeConfig) {
		return getNonInteractiveClientConfig(kubeConfig), nil
	}

	return getEmbedded()
}

func isManual(kubeConfig string) bool {
	if kubeConfig != "" {
		return true
	}
	_, inClusterErr := rest.InClusterConfig()
	return inClusterErr == nil
}

func getEmbedded() (clientcmd.ClientConfig, error) {
	return nil, fmt.Errorf("embedded only supported on linux")
}

// getNonInteractiveClientConfig builds a ClientConfig from the given
// kubeconfig path using the standard client-go loading rules.
func getNonInteractiveClientConfig(kubeConfig string) clientcmd.ClientConfig {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.DefaultClientConfig = &clientcmd.DefaultClientConfig
	if kubeConfig != "" {
		loadingRules.ExplicitPath = kubeConfig
	}
	overrides := &clientcmd.ConfigOverrides{ClusterDefaults: clientcmd.ClusterDefaults}
	return clientcmd.NewInteractiveDeferredLoadingClientConfig(loadingRules, overrides, nil)
}
