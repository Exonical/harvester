package setting

import (
	"encoding/json"

	"github.com/sirupsen/logrus"

	harvesterv1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	"github.com/harvester/harvester/pkg/settings"
	"github.com/harvester/harvester/pkg/util"
)

func (h *Handler) syncSSLParameters(setting *harvesterv1.Setting) error {
	sslParameter := &settings.SSLParameter{}
	value := setting.Value
	if value == "" {
		value = setting.Default
	}
	if err := json.Unmarshal([]byte(value), sslParameter); err != nil {
		return err
	}

	return h.updateSSLParameters(sslParameter)
}

func (h *Handler) updateSSLParameters(sslParameter *settings.SSLParameter) error {
	logrus.Infof("Update SSL Parameters: Ciphers: %s, Protocols: %s", sslParameter.Ciphers, sslParameter.Protocols)

	configMap, err := h.configmapCache.Get(util.HarvesterSystemNamespaceName, "harvester-gateway-ssl-parameters")
	if err != nil {
		return err
	}
	toUpdate := configMap.DeepCopy()
	if toUpdate.Data == nil {
		toUpdate.Data = make(map[string]string)
	}

	if sslParameter.Ciphers == "" {
		delete(toUpdate.Data, "ssl-ciphers")
	} else {
		toUpdate.Data["ssl-ciphers"] = sslParameter.Ciphers
	}
	if sslParameter.Protocols == "" {
		delete(toUpdate.Data, "ssl-protocols")
	} else {
		toUpdate.Data["ssl-protocols"] = sslParameter.Protocols
	}

	if _, err := h.configmaps.Update(toUpdate); err != nil {
		return err
	}

	return nil
}
