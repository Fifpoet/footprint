package server

import (
	"github.com/casbin/casbin/persist/file-adapter"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	Router      *gin.Engine
	FileAdapter *fileadapter.Adapter
	RedisCli    *redis.Client
}
