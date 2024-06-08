package routers

import (
	"log"
	"time"

	"github.com/FlezzProject/platform-api/internal/infrastructure/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type GatewayRouter struct {
	Base           *BaseRouter
	gatewayConfigs *config.GatewayConfigs
}

func NewGatewayRouter(gatewayConfigs *config.GatewayConfigs, baseRouter *BaseRouter) GatewayRouter {
	r := GatewayRouter{
		Base:           baseRouter,
		gatewayConfigs: gatewayConfigs,
	}
	r.Base.Gin.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		// AllowOrigins: []string{
		// 	"http://localhost:8002",
		// 	"http://localhost:3000",
		// },
		//アプリとウェブがAPIを使用できるようにするために、allow-allを設定(開発環境のみ)。
		AllowAllOrigins: true,
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"DELETE",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
			"auth-token",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))
	r.draw()
	return r
}

func (r GatewayRouter) draw() {
	log.Println("Gateway Router is drawing routes...")
	for _, route := range r.gatewayConfigs.Routes {
		log.Printf("Mapping '%v' | %v ---> %v:%v", route.Name, route.Context, route.Host, route.TargetPort)
		proxy := config.NewProxy(route)
		r.Base.Gin.Any(route.Context, func(ctx *gin.Context) {
			proxy.PassThrough(ctx)
		})
	}
}
