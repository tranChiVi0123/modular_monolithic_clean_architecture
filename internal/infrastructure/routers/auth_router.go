package routers

import (
	"github.com/FlezzProject/platform-api/internal/app/http/handler"
)

type AuthRouter struct {
	authHandler handler.AuthHandler
	Base        *BaseRouter
}

func NewAuthRouter(authHandler handler.AuthHandler, baseRouter *BaseRouter) AuthRouter {
	r := AuthRouter{
		authHandler: authHandler,
		Base:        baseRouter,
	}

	r.draw()
	return r
}

func (r AuthRouter) draw() {
	r.Base.Gin.POST("/login", r.authHandler.Login)
	r.Base.Gin.POST("/register", r.authHandler.Register)
	r.Base.Gin.POST("/logout", r.authHandler.Logout)
}
