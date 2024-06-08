package initializers

import (
	"context"

	"github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/infrastructure/config"
	"github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/infrastructure/db"
	"github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/initializers/gen"
)

func InitializeAuthService(ctx context.Context) error {
	envConfig, err := config.LoadEnvVariables(config.AUTH_SERVICE)
	if err != nil {
		return err
	}

	fdb, err := db.BuildMysqlConnection(*envConfig)
	if err != nil {
		return err
	}

	dbConfig := db.DbConfig{
		FDB: fdb,
	}

	router := gen.InitializeAuthRouting(dbConfig, envConfig.SecretKey)

	router.Base.Run(envConfig.Route.Host, envConfig.Route.TargetPort)

	return nil
}
