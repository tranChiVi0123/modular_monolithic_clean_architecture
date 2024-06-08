package routers

import (
	"github.com/FlezzProject/platform-api/internal/app/http/handler"
	"github.com/FlezzProject/platform-api/internal/app/http/middleware"
)

type UserRouter struct {
	userHandler    handler.UserHandler
	authMiddleware middleware.AuthMiddleware
	Base           *BaseRouter
}

func NewUserRouter(
	userHandler handler.UserHandler,
	baseRouter *BaseRouter,
	authMiddleware middleware.AuthMiddleware,
) UserRouter {
	r := UserRouter{
		userHandler:    userHandler,
		Base:           baseRouter,
		authMiddleware: authMiddleware,
	}

	r.draw()
	return r
}

func (r *UserRouter) draw() {
	authRequirePath := r.Base.Gin.Group("/")
	authRequirePath.Use(r.authMiddleware.Execute)
	{
		authRequirePath.GET("/example", r.userHandler.Show)
	}
}
