package routers

import (
	"log"

	"github.com/FlezzProject/platform-api/internal/app/http/handler"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct {
	Gin            *gin.Engine
	healthzHandler *handler.HealthzHandler
}

func NewBaseRouter(healthzHandler handler.HealthzHandler) *BaseRouter {
	g := gin.New()
	r := &BaseRouter{
		Gin:            g,
		healthzHandler: &healthzHandler,
	}
	r.Gin.Use(gin.Logger())
	r.draw()
	return r
}

func (r *BaseRouter) draw() {
	// Define routes here
}

func (r *BaseRouter) Run(serviceName string, port string) {
	log.Println(serviceName + " Service is running on port " + port + "...")
	r.Gin.Run(":" + port)
}
