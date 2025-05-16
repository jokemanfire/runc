//go:build linux && loong64

package userns

import "github.com/opencontainers/runc/libcontainer/configs"

func GetUserNamespaceMappings(nsPath string) (uidMap, gidMap []configs.IDMap, err error) {
	return nil, nil, nil
}

func IsSameMapping(a, b []configs.IDMap) bool {
	return true
}
