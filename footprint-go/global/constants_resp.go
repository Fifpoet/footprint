package global

import (
	"github.com/gin-gonic/gin"
)

// 用户模块
const (
	CodeSuccess      = "20000"
	CodeBadReq       = "40000"
	CodeUnAuth       = "40003"
	CodeNoUser       = "40006"
	CodeAuthError    = "40007"
	CodeTokenInvalid = "40008"
	CodeDaoError     = "50001"

	MsgSuccess      = "success"
	MsgBadReq       = "invalid params"
	MsgUnAuth       = "login please"
	MsgNoUser       = "no user found"
	MsgAuthError    = "auth error"
	MsgTokenInvalid = "token invalid"
	MsgDaoError     = "error operating database"
)

// 文件上传模块
const (
	CodeNoFileRecv      = "51000"
	CodeFileCantOpen    = "51001"
	CodeFileUploadError = "51002"

	MsgNoFileRecv      = "can't receive file from req"
	MsgFileCantOpen    = "file cannot open"
	MsgFileUploadError = "file upload failed"
)

func Resp(c *gin.Context, code string, msg string, httpCode int, err error) {
	var errStack string
	if err != nil {
		errStack = msg + ": " + err.Error()
	} else {
		errStack = msg
	}
	c.JSON(httpCode, gin.H{
		"code": code,
		"msg":  errStack,
	})
}
