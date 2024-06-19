package controllers

import (
	"dxj/internal/models"
	"dxj/internal/pkg/e"
	"dxj/internal/services"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticlesController struct {
}

var articleService = services.ArticleService{}

// 新建文章
func (ArticlesController) CreateArticle(c *gin.Context) {
	var data models.Article
	c.BindJSON(&data)
	article := articleService.CreateArticle(&data)
	e.SendData(c, e.SUCCESS, article)
}

// 通过文章id获取文章详情
func (ArticlesController) GetArticleById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		// 处理转换错误，例如返回错误响应或日志记录
		log.Printf("Error converting id to int64: %v", err)
	}
	article := articleService.GetArticleById(idInt)
	e.SendData(c, e.SUCCESS, article)
}

// 分页获取文章详情列表
func (ArticlesController) GetArticleList(c *gin.Context) {
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	pageSize, _ := strconv.ParseInt(c.Query("pageSize"), 10, 64)
	articles, total := articleService.GetArticleList(page, pageSize)
	e.SendData(c, e.SUCCESS, gin.H{
		"list":  articles,
		"total": total,
		"page":  page,
	},
	)
}
