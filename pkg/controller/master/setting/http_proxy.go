package setting

import (
	"encoding/json"

	harvesterv1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	"github.com/harvester/harvester/pkg/util"
)

const (
	fleetLocalNamespace = "fleet-local"
	localClusterName    = "local"
)

func (h *Handler) syncHTTPProxy(setting *harvesterv1.Setting) error {
	var httpProxyConfig util.HTTPProxyConfig
	value := setting.Value
	if value == "" {
		value = setting.Default
		if value == "" {
			value = "{}"
		}
	}
	if err := json.Unmarshal([]byte(value), &httpProxyConfig); err != nil {
		return err
	}
	backupConfig := map[string]string{
		util.HTTPProxyEnv:  httpProxyConfig.HTTPProxy,
		util.HTTPSProxyEnv: httpProxyConfig.HTTPSProxy,
		util.NoProxyEnv:    util.AddBuiltInNoProxy(httpProxyConfig.NoProxy),
	}
	if err := h.updateBackupSecret(backupConfig); err != nil {
		return err
	}
	if err := h.syncClusterHTTPProxy(httpProxyConfig); err != nil {
		return err
	}

	// Redeploy system services. The proxy envs will be injected by the mutation webhook.
	if err := h.redeployDeployment(util.CattleSystemNamespaceName, "rancher"); err != nil {
		return err
	}
	return h.redeployDeployment(h.namespace, "harvester")
}

func (h *Handler) syncClusterHTTPProxy(httpProxyConfig util.HTTPProxyConfig) error {
	proxyConfigMap, err := h.configmapCache.Get(util.HarvesterSystemNamespaceName, "harvester-http-proxy")
	if err != nil {
		return err
	}
	toUpdate := proxyConfigMap.DeepCopy()
	if toUpdate.Data == nil {
		toUpdate.Data = make(map[string]string)
	}
	toUpdate.Data[util.HTTPProxyEnv] = httpProxyConfig.HTTPProxy
	toUpdate.Data[util.HTTPSProxyEnv] = httpProxyConfig.HTTPSProxy
	toUpdate.Data[util.NoProxyEnv] = util.AddBuiltInNoProxy(httpProxyConfig.NoProxy)
	_, err = h.configmaps.Update(toUpdate)
	return err
}
