package api

import (
	"github.com/fifpoet/footprint/dao"
	"github.com/fifpoet/footprint/global"
	"github.com/fifpoet/footprint/model"
	utils "github.com/fifpoet/footprint/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func AddArticle(c *gin.Context) {
	req := &model.AddArticleReq{}
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
		Lat:        req.Lat,
		Lng:        req.Lng,
		City:       req.City,
		Model:      gorm.Model{},
	}
	err = dao.ArticleRepo{}.Add(article)
	if err != nil {
		global.Resp(c, global.DaoError, nil)
		return
	}
	global.Resp(c, global.Success, nil)
}

func GetArticle(c *gin.Context) {
	pageNo, _ := utils.String2Int(c.Query("pageNo"))
	pageSize, _ := utils.String2Int(c.Query("pageSize"))
	arts, err := dao.ArticleRepo{}.FindByPage(pageNo, pageSize)
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

func UpdateArticle(c *gin.Context) {
	req := &model.UpdateArticleReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Resp(c, global.BadReq, err)
		return
	}
	err = dao.ArticleRepo{}.Update(model.Article{
		UserId:     req.UserId,
		CategoryId: req.CategoryId,
		Cover:      req.Cover,
		Title:      req.Title,
		Content:    req.Content,
		Lat:        req.Lat,
		Lng:        req.Lng,
		City:       req.City,
		Model: gorm.Model{
			ID:        req.ID,
			UpdatedAt: time.Now(),
		},
	})
	if err != nil {
		global.Resp(c, global.DaoError, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"msg":  "success",
	})
}

func DelArticleById(c *gin.Context) {
	id := c.Param("id")
	err := dao.ArticleRepo{}.DelById(id)
	if err != nil {
		global.Resp(c, global.DaoError, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"msg":  "success",
	})
}
