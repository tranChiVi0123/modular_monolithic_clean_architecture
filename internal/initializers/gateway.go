package initializers

import (
	"context"

	"github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/infrastructure/config"
	"github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/initializers/gen"
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
