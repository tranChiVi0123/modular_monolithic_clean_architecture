//go:build wireinject
// +build wireinject

package gen

import (
	"github.com/FlezzProject/platform-api/internal/app/http/handler"
	"github.com/FlezzProject/platform-api/internal/app/http/middleware"
	"github.com/FlezzProject/platform-api/internal/domain/irepository"
	"github.com/FlezzProject/platform-api/internal/infrastructure/config"
	"github.com/FlezzProject/platform-api/internal/infrastructure/db"
	"github.com/FlezzProject/platform-api/internal/infrastructure/iusecase"
	"github.com/FlezzProject/platform-api/internal/infrastructure/routers"
	hrp "github.com/FlezzProject/platform-api/internal/modules/healthz/repository"
	huc "github.com/FlezzProject/platform-api/internal/modules/healthz/usecase"
	urp "github.com/FlezzProject/platform-api/internal/modules/user/repository"
	uuc "github.com/FlezzProject/platform-api/internal/modules/user/usecase"
	"github.com/google/wire"
)

func InitializeGatewayRouting(gatewayConfigs *config.GatewayConfigs) routers.GatewayRouter {
	wire.Build(
		routers.NewGatewayRouter,
		routers.NewBaseRouter,

		// handler
		handler.NewHealthzHandler,

		// usecases
		huc.NewHealthzUsecase,
		wire.Bind(new(iusecase.IHealthzUsecase), new(huc.HealthzUsecase)),

		// repositories
		hrp.NewHealthzRepository,
		wire.Bind(new(irepository.IHealthzRepository), new(hrp.HealthzRepository)),
	)
	return routers.GatewayRouter{}
}

func InitializeAuthRouting(dbConfig db.DbConfig, secretKey string) routers.AuthRouter {
	wire.Build(
		routers.NewAuthRouter,
		routers.NewBaseRouter,

		// handlers
		handler.NewAuthHandler,
		handler.NewHealthzHandler,

		// usecases
		uuc.NewUserUsecase,
		wire.Bind(new(iusecase.IUserUsecase), new(uuc.UserUsecase)),
		huc.NewHealthzUsecase,
		wire.Bind(new(iusecase.IHealthzUsecase), new(huc.HealthzUsecase)),

		// repositories
		urp.NewUserRepository,
		wire.Bind(new(irepository.IUserRepository), new(*urp.UserRepository)),
		hrp.NewHealthzRepository,
		wire.Bind(new(irepository.IHealthzRepository), new(hrp.HealthzRepository)),
	)
	return routers.AuthRouter{}
}

func InitializeUserRouting(dbConfig db.DbConfig, secretKey string) routers.UserRouter {
  wire.Build(
    routers.NewUserRouter,
    routers.NewBaseRouter,

    // handlers
    handler.NewUserHandler,
    handler.NewHealthzHandler,

    // middlewares
    middleware.NewAuthMiddleware,

    // usecases
    uuc.NewUserUsecase,
    wire.Bind(new(iusecase.IUserUsecase), new(uuc.UserUsecase)),
    huc.NewHealthzUsecase,
    wire.Bind(new(iusecase.IHealthzUsecase), new(huc.HealthzUsecase)),

    // repositories
    urp.NewUserRepository,
    wire.Bind(new(irepository.IUserRepository), new(*urp.UserRepository)),
    hrp.NewHealthzRepository,
    wire.Bind(new(irepository.IHealthzRepository), new(hrp.HealthzRepository)),
  )
  return routers.UserRouter{}
}
