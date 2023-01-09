// Package bootstrap
package bootstrap

import (
	"cake-store/cake-store/internal/consts"
	"cake-store/cake-store/pkg/logger"
	"cake-store/cake-store/pkg/msgx"
)

func RegistryMessage()  {
	err := msgx.Setup("msg.yaml", consts.ConfigPath)
	if err != nil {
		logger.Fatal(logger.MessageFormat("file message multi language load error %s", err.Error()))
	}

}
