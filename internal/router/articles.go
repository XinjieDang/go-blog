package router

import (
	"dxj/internal/controllers"
	"dxj/internal/global"

	"github.com/gin-gonic/gin"
)

func LoadAricleRoutes(r *gin.Engine) *gin.RouterGroup {
	article := r.Group(global.Config.Api.Prefix + "/article")
	categoryController := controllers.ArticlesController{}
	{
		article.POST("/create", categoryController.CreateArticle)
		article.GET("/articleInfo/:id", categoryController.GetArticleById)
		article.GET("/list", categoryController.GetArticleList)
	}
	return article
}
