package service

import (
	"github.com/gin-gonic/gin"
	"api_gateway/internal/utils"
	"log"
    "net/http/httputil"
    "net/url"
)

func Run(){
	utils.LogInit()

	routesRun()
}

func reverseProxy(target string) gin.HandlerFunc {
    return func(c *gin.Context) {
        remote, err := url.Parse(target)
        if err != nil {
			respondError(c, 502, "invalid target")
            return
        }

        proxy := httputil.NewSingleHostReverseProxy(remote)
        proxy.ServeHTTP(c.Writer, c.Request)
    }
}

func routesRun(){
	cfg, err := utils.GetConfig()
	if err != nil{
		log.Println(err)
	}

	var r = gin.Default()

	//address service
	r.Any("/api/suggest", reverseProxy(cfg.AddressService.URL()))
	r.Any("/api/distance", reverseProxy(cfg.AddressService.URL()))
	r.Any("/api/address", reverseProxy(cfg.AddressService.URL()))
	r.Any("/api/address/*proxyPath", reverseProxy(cfg.AddressService.URL()))

	//user service
	r.Any("/api/roles", reverseProxy(cfg.UserService.URL()))
	r.Any("/api/users", reverseProxy(cfg.UserService.URL()))
	r.Any("/api/roles/*proxyPath", reverseProxy(cfg.UserService.URL()))
	r.Any("/api/users/*proxyPath", reverseProxy(cfg.UserService.URL()))
	r.Any("/api/login", reverseProxy(cfg.UserService.URL()))
	r.Any("/api/logout", reverseProxy(cfg.UserService.URL()))
	r.Any("/api/verify", reverseProxy(cfg.UserService.URL()))

	//product service
	r.Any("/api/units", reverseProxy(cfg.ProductService.URL()))
	r.Any("/api/products", reverseProxy(cfg.ProductService.URL()))
	r.Any("/api/units/*proxyPath", reverseProxy(cfg.ProductService.URL()))
	r.Any("/api/products/*proxyPath", reverseProxy(cfg.ProductService.URL()))

	r.Run(cfg.Service.EndPoint())
}
