package config

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	errors_handler "github.com/tranChiVi0123/modular_monolithic_clean_architecture/pkg/common/errors"
)

type Proxy struct {
	Route Route
}

func NewProxy(route Route) Proxy {
	return Proxy{
		Route: route,
	}
}

func (p Proxy) PassThrough(ctx *gin.Context) {
	urlStr := p.Route.Protocol + p.Route.Host + ":" + p.Route.TargetPort
	targetURL, err := url.Parse(urlStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errors_handler.New500ErrorResponse(err))
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	ctx.Request.URL.Path = ctx.Param("proxyPath")
	ctx.Request.Host = p.Route.Host
	ctx.Request.Header.Set("X-PROXY", "FLEZZ_GATEWAY")

	log.Printf("Proxying to: %v%v", urlStr, ctx.Request.URL.Path)

	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}
