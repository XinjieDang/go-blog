package router

import (
	"dxj/internal/controllers"
	"dxj/internal/global"
	"dxj/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func LoadUserRoutes(r *gin.Engine) *gin.RouterGroup {
	user := r.Group(global.Config.Api.Prefix + "/user")
	//私有路由添加中间件
	user.Use(middlewares.VerifyJWTUser())
	userController := controllers.UserController{}
	{
		user.GET("/getHello", userController.GetUserInfo)
		user.GET("/getUserInfoById/:id", userController.GetUserInfoById)
	}
	return user
}
