package middleware

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/persist"
	"github.com/fifpoet/footprint/global"
	"github.com/fifpoet/footprint/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		mc := utils.ExtractTokenPayload(c.Request)
		if mc == nil {
			global.Resp(c, global.UnAuth, nil)
			c.Abort()
			return
		}
		if mc.ExpiresAt.Before(time.Now()) {
			global.Resp(c, global.TokenExpired, nil)
			c.Abort()
			return
		}
		// 单次请求的上下文缓存
		c.Set("userId", mc.UserId)
		c.Set("userName", mc.UserName)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

// Authorize determines if current subject has been authorized to take an action on an object.
func Authorize(obj string, act string, adapter persist.Adapter) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Get("userId")
		username, _ := c.Get("userName")
		if id == "" {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}
		// casbin enforces policy
		ok, err := enforce(username.(string), obj, act, adapter)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(500, "error occurred when authorizing user: "+err.Error())
			return
		}
		if !ok {
			c.AbortWithStatusJSON(403, "forbidden")
			return
		}
		c.Next()
	}
}

func enforce(sub string, obj string, act string, adapter persist.Adapter) (bool, error) {
	//TODO RBAC权限入库
	enforcer := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	err := enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}
	ok := enforcer.Enforce(sub, obj, act)
	return ok, nil
}
