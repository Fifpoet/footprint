package api

import (
	"github.com/fifpoet/footprint/dao"
	"github.com/fifpoet/footprint/global"
	"github.com/fifpoet/footprint/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func UploadArticle(c *gin.Context) {
	req := &model.UploadArticleReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Resp(c, global.CodeBadReq, global.MsgBadReq, http.StatusBadRequest, err)
		return
	}
	userId, _ := c.Get("userId")
	article := model.Article{
		UserId:     userId.(uint),
		CategoryId: 0,
		Cover:      req.Cover,
		Title:      req.Title,
		Content:    req.Detail,
		Lat:        req.Lat,
		Lng:        req.Lng,
		City:       req.Location.City,
		District:   req.Location.District,
		Province:   req.Location.Province,
		Model:      gorm.Model{},
	}
	err = dao.ArticleRepo{}.Update(article)
	if err != nil {
		global.Resp(c, global.CodeDaoError, global.MsgDaoError, http.StatusInternalServerError, nil)
		return
	}
	global.Resp(c, global.CodeSuccess, global.MsgSuccess, http.StatusOK, nil)
}

func GetArticle(c *gin.Context) {
	var id uint64
	id, _ = strconv.ParseUint(c.Query("id"), 10, 64)
	arts, err := dao.ArticleRepo{}.FindById(uint(id))
	if err != nil {
		global.Resp(c, global.CodeDaoError, global.MsgDaoError, http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     20000,
		"msg":      "success",
		"articles": arts,
	})
}
