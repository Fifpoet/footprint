package api

import (
	"github.com/fifpoet/footprint/global"
	"github.com/fifpoet/footprint/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		global.Resp(c, global.CodeNoFileRecv, global.MsgNoFileRecv, http.StatusInternalServerError, err)
		return
	}
	//Reader流打开文件
	reader, err := file.Open()
	if err != nil {
		global.Resp(c, global.CodeFileCantOpen, global.MsgFileCantOpen, http.StatusInternalServerError, err)
		return
	}
	defer reader.Close()
	//上传
	conf := global.FP_CONFIG.Minio
	url, err := service.UploadFile(global.FP_MINIO, conf.Bucket, conf.Location, file.Filename, reader, file.Size)
	if err != nil {
		global.Resp(c, global.CodeFileUploadError, global.MsgFileUploadError, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"msg":  "success",
		"size": file.Size,
		"link": url,
	})
}
