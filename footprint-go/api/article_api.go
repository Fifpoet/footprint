package api

import (
	"github.com/fifpoet/footprint/dao"
	"github.com/fifpoet/footprint/global"
	"github.com/fifpoet/footprint/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func UploadArticle(c *gin.Context) {
	req := &model.UploadArticleReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Resp(c, global.BadReq, err)
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
		global.Resp(c, global.DaoError, nil)
		return
	}
	global.Resp(c, global.Success, nil)
}

func GetArticle(c *gin.Context) {
	userName, _ := c.Get("userName")
	user, _ := dao.FindByName(userName.(string))
	arts, err := dao.ArticleRepo{}.FindById(user.ID)
	if err != nil {
		global.Resp(c, global.DaoError, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     20000,
		"msg":      "success",
		"articles": arts,
	})
}
