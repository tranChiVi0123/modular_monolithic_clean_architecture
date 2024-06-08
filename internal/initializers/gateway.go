package initializers

import (
	"context"

	"github.com/FlezzProject/platform-api/internal/infrastructure/config"
	"github.com/FlezzProject/platform-api/internal/initializers/gen"
)

func InitializeGateway(ctx context.Context) error {
	gateWayConfigs, err := config.LoadGatewayConfigs()

	if err != nil {
		return err
	}

	router := gen.InitializeGatewayRouting(gateWayConfigs)

	router.Base.Run(gateWayConfigs.Host, gateWayConfigs.GatewayPort)
	return nil
}
