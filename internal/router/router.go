package router

import (
	"dxj/internal/global"
	"dxj/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	gin.SetMode(global.Config.Server.GinMode)
	Router := gin.Default()
	Router.Group("/api/v1")
	// Global middlewares
	Router.Use(middlewares.ErrorHandle())
	Router.Use(middlewares.Cors())
	// public routes, no auth required
	LoadPublicRoutes(Router)
	// user routes
	LoadUserRoutes(Router)
	LoadCategoryRoutes(Router)
	LoadAricleRoutes(Router)
	return Router
}
