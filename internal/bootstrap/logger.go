// Package bootstrap
package bootstrap

import (
	"cake-store/cake-store/internal/appctx"
	"cake-store/cake-store/pkg/logger"
	"cake-store/cake-store/pkg/util"
)

func RegistryLogger(cfg *appctx.Config) {
	logger.Setup(logger.Config{
		Environment: util.EnvironmentTransform(cfg.App.Env),
		Debug:       cfg.App.Debug,
		Level:       cfg.Logger.Level,
		ServiceName: cfg.Logger.Name,
	})
}

