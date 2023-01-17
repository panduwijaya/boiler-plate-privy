package http

import (
	"context"

	"cake-store/cake-store/internal/consts"
	"cake-store/cake-store/internal/handler"
	"cake-store/cake-store/internal/server"
	"cake-store/cake-store/pkg/logger"
)

// Start function handler starting http listener
func Start(ctx context.Context) {

	serve := server.NewHTTPServer()
	defer serve.Done()
	logger.Info(logger.MessageFormat("starting cake-store services... %d", serve.Config().App.Port), logger.EventName(consts.LogEventNameServiceStarting))

	//starting subscribe pubsub
	handler.LISTEN()

	if err := serve.Run(ctx); err != nil {
		logger.Warn(logger.MessageFormat("service stopped, err:%s", err.Error()), logger.EventName(consts.LogEventNameServiceStarting))
	}

	return
}
