package setting

import (
	"github.com/sirupsen/logrus"

	harvesterv1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
)

func (h *Handler) syncKubeconfigTTL(setting *harvesterv1.Setting) error {
	logrus.Infof("Kubeconfig TTL setting updated to %s (applied via Harvester settings)", setting.Value)
	return nil
}
