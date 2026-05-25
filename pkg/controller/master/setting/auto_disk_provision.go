package setting

import (
	"github.com/sirupsen/logrus"

	harvesterv1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
)

func (h *Handler) syncNDMAutoProvisionPaths(setting *harvesterv1.Setting) error {
	logrus.Infof("NDM auto-provision paths updated to: %s (apply via Helm chart values)", setting.Value)
	return nil
}
