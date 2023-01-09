// Package bootstrap
package bootstrap

import (
	"fmt"

	"cake-store/cake-store/internal/common"
	"cake-store/cake-store/pkg/generator"
	"cake-store/cake-store/pkg/logger"
)

// RegistrySnowflake setup snowflake generator
func RegistrySnowflake()  {
	hs := common.GetHostname()
	nodeID := uint64(common.GetNodeID(hs))

	lf := logger.NewFields(
		logger.EventName("SetupSnowflake"),
		logger.Any("node_id", nodeID),
		logger.Any("hostname", hs),
		)

	logger.Info(fmt.Sprintf(`generate node id for snowflake is %d`, nodeID), lf...)
	generator.Setup(nodeID)
}
