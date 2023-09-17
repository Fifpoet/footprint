package jwt

import (
	"github.com/fifpoet/footprint/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthLogin 读取user校验 签发token
func AuthLogin(c *gin.Context) {
	user := &model.UserInfo{}
	// 直接绑定到user结构体
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 40000,
			"msg":  "invalid param",
		})
		return
	}
	// TODO 读取db校验密码

	token, _ := GenerateToken(*user)
	c.JSON(http.StatusBadRequest, gin.H{
		"code":  20000,
		"msg":   "success",
		"token": token,
	})
}

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.Split(authHeader, ".")
		if len(parts) != 3 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		mc, err := ParseToken(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		m := mc.User
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", m.Name)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
