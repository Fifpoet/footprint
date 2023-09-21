package server

import (
	"github.com/fifpoet/footprint/api"
	"github.com/fifpoet/footprint/middleware"
	"github.com/gin-gonic/gin"
)

func (s *Server) InitRoutes() {

	s.Router.POST("/login", api.Login)
	//为路由组添加log和panic恢复的中间件
	authorized := s.Router.Group("/")
	authorized.Use(gin.Logger())
	authorized.Use(gin.Recovery())
	authorized.Use(middleware.TokenAuthMiddleware())
	{
		authorized.POST("/api/todo", middleware.Authorize("resource", "write", s.FileAdapter), api.Logout)
		authorized.GET("/api/todo", middleware.Authorize("resource", "read", s.FileAdapter), api.Logout)
		authorized.POST("/logout", api.Logout)
		authorized.POST("/refresh", api.Refresh)
	}

}
