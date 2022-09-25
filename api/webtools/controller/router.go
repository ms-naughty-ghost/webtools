package controller

import (
	"net/http"
	"time"
	"webtools/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"*",
		},
		AllowMethods: []string{
			"GET",
			"HEAD",
		},
		AllowHeaders: []string{
			"*",
		},
		MaxAge: 12 * time.Hour,
	}))
	for _, route := range routes {
		switch route.Method {
		case http.MethodHead:
			router.HEAD(route.Pattern, route.HandlerFunc)
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}
	return router
}

var routes = []Route{
	{
		"死活監視",
		http.MethodGet,
		"/v1/ping",
		func(c *gin.Context) {},
	},
	{
		"セキュアなランダム文字列生成",
		http.MethodGet,
		"/v1/securerandomstring",
		service.SecureRandomStr,
	},
}
