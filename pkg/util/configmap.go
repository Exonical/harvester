package util

import "github.com/harvester/harvester/pkg/util/name"

func GetRestoreVMConfigMapName(upgradeName string) string {
	return name.SafeConcatName(upgradeName, RestoreVMConfigMap)
}
