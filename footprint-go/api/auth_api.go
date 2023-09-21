package api

import (
	"github.com/fifpoet/footprint/dao"
	"github.com/fifpoet/footprint/model"
	"github.com/fifpoet/footprint/service/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login 读取user校验 签发token
func Login(c *gin.Context) {
	user := &model.UserInfo{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 40000,
			"msg":  "invalid param",
		})
		return
	}
	// TODO 读取db校验密码
	dbUser, err := dao.UserRepo{}.FindById(user.Model.ID)
	if &dbUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 40008,
			"msg":  "user not found",
		})
		return
	}
	if dbUser.Password != user.Password || dbUser.Name != user.Name {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 40003,
			"msg":  "name or password error",
		})
		return
	}
	// 双token
	access, refresh, _ := jwt.GenerateToken(*user)
	c.JSON(http.StatusOK, gin.H{
		"code":          20000,
		"msg":           "success",
		"access_token":  access,
		"refresh_token": refresh,
	})
}

// Refresh access过期 Refresh没过期则刷新
// TODO
func Refresh(c *gin.Context) {
	tks := map[string]string{}
	if err := c.ShouldBindJSON(&tks); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	//校验token
}

func Logout(c *gin.Context) {
	//TODO
}