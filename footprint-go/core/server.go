package core

import (
	"fmt"
	fileadapter "github.com/casbin/casbin/persist/file-adapter"
	"github.com/fifpoet/footprint/api"
	"github.com/fifpoet/footprint/global"
	"github.com/fifpoet/footprint/middleware"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	InitRoutes(r)
	address := fmt.Sprintf(":%d", global.FP_CONFIG.App.Port)
	s := InitializeServer(address, r)

	global.FP_LOG.Info("server run success on ", zap.String("address", address))
	s.ListenAndServe()
}

/*
这些配置可以帮助我们优化HTTP服务器的性能和安全性。通过设置超时时间和最大字节数等参数，可以防止一些潜在的安全问题和性能问题。
例如，设置超时时间可以防止客户端故意保持连接而导致的资源浪费，设置最大字节数可以防止客户端发送过大的请求头而导致的资源浪费和安全问题。
*/
func InitializeServer(address string, router *gin.Engine) server {
	// 使用endless库创建一个HTTP服务器，其中address是服务器的监听地址（如:8080），router是HTTP请求路由器。
	s := endless.NewServer(address, router)

	// 设置HTTP请求头的读取超时时间为20秒，如果在20秒内未读取到请求头，则会返回一个超时错误。
	s.ReadHeaderTimeout = 20 * time.Second

	// 设置HTTP响应体的写入超时时间为20秒，如果在20秒内未将响应体写入完成，则会返回一个超时错误。
	s.WriteTimeout = 20 * time.Second

	// 设置HTTP请求头的最大字节数为1MB。如果请求头超过1MB，则会返回一个错误。
	s.MaxHeaderBytes = 1 << 20

	return s
}

func InitRoutes(router *gin.Engine) {
	f := fileadapter.NewAdapter("config/basic_policy.csv")

	router.Use(middleware.CorsAllowAll())
	router.POST("/login", api.Login)
	//为路由组添加log和panic恢复的中间件
	authorized := router.Group("/")
	authorized.Use(gin.Logger())
	authorized.Use(gin.Recovery())
	authorized.Use(middleware.TokenAuthMiddleware())
	authorized.Use(middleware.CorsAllowAll())
	{
		authorized.POST("/api/article", middleware.Authorize("resource", "write", f), api.UploadArticle)
		authorized.GET("/api/article", middleware.Authorize("resource", "read", f), api.GetArticle)
		authorized.POST("/logout", api.Logout)
		authorized.POST("/refresh", api.Refresh)
	}

}
