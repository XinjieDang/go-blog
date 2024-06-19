package router

import (
	"dxj/internal/controllers"
	"dxj/internal/global"
	"dxj/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func LoadCategoryRoutes(r *gin.Engine) *gin.RouterGroup {
	category := r.Group(global.Config.Api.Prefix + "/category")
	category.Use(middlewares.VerifyJWTUser())
	categoryController := controllers.CategoryController{}
	{
		category.GET("/list", categoryController.GetCategoryList)
		category.POST("/create", categoryController.CreateCategory)
	}
	return category
}
