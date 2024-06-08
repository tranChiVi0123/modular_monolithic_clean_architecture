package initializers

import (
	"context"

	"github.com/FlezzProject/platform-api/internal/infrastructure/config"
	"github.com/FlezzProject/platform-api/internal/infrastructure/db"
	"github.com/FlezzProject/platform-api/internal/initializers/gen"
)

func InitializeUserService(ctx context.Context) error {
	envConfig, err := config.LoadEnvVariables(config.USER_SERVICE)
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

	router := gen.InitializeUserRouting(dbConfig, envConfig.SecretKey)

	router.Base.Run(envConfig.Route.Host, envConfig.Route.TargetPort)

	return nil
}
