package api

import (
	"github.com/fifpoet/footprint/dao"
	"github.com/fifpoet/footprint/global"
	"github.com/fifpoet/footprint/model"
	"github.com/fifpoet/footprint/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// Login 校验user; 签发token; 返回用户信息
func Login(c *gin.Context) {
	user := &model.LoginReq{}
	err := c.ShouldBindJSON(&user)
	if err != nil || user == nil {
		global.Resp(c, global.BadReq, err)
		return
	}
	dbUser, err := dao.FindByName(user.UserName)
	if dbUser == nil {
		global.Resp(c, global.NoUser, err)
		return
	}
	if dbUser.Password != user.Password {
		global.Resp(c, global.AuthError, nil)
		return
	}
	// 双token
	access, refresh, err := utils.GenerateToken(*dbUser)
	c.JSON(http.StatusOK, gin.H{
		"code":          20000,
		"msg":           "success",
		"access_token":  access,
		"refresh_token": refresh,
		// 嵌套结构体 不自动转小写
		"user": model.LoginVO{
			Name:   dbUser.Name,
			Email:  dbUser.Email,
			Avatar: dbUser.Avatar,
			Role:   utils.DotString2Arr(dbUser.Role),
			ID:     dbUser.Model.ID,
		},
	})
}

// Refresh 请求传入refresh token, context中已经有userName, 直接刷新两个token
func Refresh(c *gin.Context) {
	name, _ := c.Get("userName")
	id, _ := c.Get("userId")
	access, refresh, err := utils.GenerateToken(model.User{
		Name:  name.(string),
		Model: gorm.Model{ID: id.(uint)},
	})
	if err != nil {
		global.Resp(c, global.InternalError, nil)
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
		global.Resp(c, global.BadReq, err)
		return
	}
	err = dao.Add(*user)
	if err != nil {
		global.Resp(c, global.DaoError, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"msg":  "success",
	})
}
