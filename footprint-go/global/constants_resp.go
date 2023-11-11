package global

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RespObj struct {
	code  int
	hcode int
	msg   string
	meta  map[string]string
}

// 用户模块
var (
	Success       = RespObj{20000, http.StatusOK, "", nil}
	BadReq        = RespObj{40000, http.StatusBadRequest, "invalid params", nil}
	UnAuth        = RespObj{40003, http.StatusForbidden, "login please", nil}
	NoUser        = RespObj{40006, http.StatusBadRequest, "no user found", nil}
	AuthError     = RespObj{40007, http.StatusBadRequest, "auth error", nil}
	TokenExpired  = RespObj{40008, http.StatusForbidden, "token invalid", nil}
	DaoError      = RespObj{50001, http.StatusInternalServerError, "error operating database", nil}
	InternalError = RespObj{50002, http.StatusInternalServerError, "internal error", nil}

	NoFileRecv      = RespObj{51000, http.StatusInternalServerError, "can't receive file from req", nil}
	FileCantOpen    = RespObj{51000, http.StatusInternalServerError, "file cannot open", nil}
	FileUploadError = RespObj{51000, http.StatusInternalServerError, "file upload failed", nil}
)

func Resp(c *gin.Context, r RespObj, err error) {
	var errStack string
	if err != nil {
		errStack = r.msg + ": " + err.Error()
	} else {
		errStack = r.msg
	}
	c.JSON(r.hcode, gin.H{
		"code": r.code,
		"msg":  errStack,
	})
}
