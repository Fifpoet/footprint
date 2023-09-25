package global

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CodeSuccess = "20000"
	CodeBadReq  = "40000"
	CodeUnAuth  = "40003"
	CodeNoUser  = "40006"
)

const (
	MsgSuccess = "success"
	MsgBadReq  = "invalid params"
	MsgUnAuth  = "login please"
	MsgNoUser  = "no user found"
)

func Resp(c *gin.Context, code string, msg string, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": code,
		"msg":  msg + err.Error(),
	})
}
