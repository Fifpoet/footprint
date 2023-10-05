package api

import (
	"github.com/fifpoet/footprint/dao"
	"github.com/fifpoet/footprint/global"
	"github.com/fifpoet/footprint/model"
	"github.com/fifpoet/footprint/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// Login 读取user校验 签发token
func Login(c *gin.Context) {
	user := &model.LoginReq{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		global.Resp(c, global.CodeBadReq, global.MsgBadReq, http.StatusBadRequest, err)
		return
	}
	dbUser, err := dao.FindByName(user.UserName)
	if &dbUser == nil {
		global.Resp(c, global.CodeNoUser, global.MsgNoUser, http.StatusUnauthorized, err)
		return
	}
	if dbUser.Password != user.Password || dbUser.Name != user.UserName {
		global.Resp(c, global.CodeAuthError, global.MsgAuthError, http.StatusUnauthorized, nil)
		return
	}
	// 双token
	access, refresh, err := service.GenerateToken(*dbUser)
	c.JSON(http.StatusOK, gin.H{
		"code":          20000,
		"msg":           "success",
		"access_token":  access,
		"refresh_token": refresh,
	})
}

// Refresh 请求传入refresh token, 直接刷新两个token
func Refresh(c *gin.Context) {
	name, _ := c.Get("userName")
	id, _ := c.Get("userId")
	access, refresh, err := service.GenerateToken(model.User{
		Name:  name.(string),
		Model: gorm.Model{ID: id.(uint)},
	})
	if err != nil {
		global.Resp(c, global.CodeInternalError, global.MsgInternalError, http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":          20000,
		"msg":           "success",
		"access_token":  access,
		"refresh_token": refresh,
	})
}

func Register(c *gin.Context) {
	user := &model.RegisterReq{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		global.Resp(c, global.CodeBadReq, global.MsgBadReq, http.StatusBadRequest, err)
		return
	}
	err = dao.Add(*user)
	if err != nil {
		global.Resp(c, global.CodeDaoError, global.MsgDaoError, http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"msg":  "success",
	})
}
