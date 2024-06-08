//go:build wireinject
// +build wireinject

package gen

import (
	"github.com/google/wire"
	"github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/app/http/handler"
	"github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/app/http/middleware"
	"github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/domain/irepository"
	"github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/infrastructure/config"
	"github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/infrastructure/db"
	"github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/infrastructure/iusecase"
	"github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/infrastructure/routers"
	hrp "github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/modules/healthz/repository"
	huc "github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/modules/healthz/usecase"
	urp "github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/modules/user/repository"
	uuc "github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/modules/user/usecase"
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
