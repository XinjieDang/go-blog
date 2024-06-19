package router

import (
	"dxj/internal/controllers"
	"dxj/internal/global"

	"github.com/gin-gonic/gin"
)

func LoadPublicRoutes(r *gin.Engine) *gin.RouterGroup {

	public := r.Group(global.Config.Api.Prefix + "/public")
	publicController := controllers.PublicController{}
	{
		public.GET("/getHello", publicController.GetHello)
		public.POST("/login", publicController.Login)
		public.POST("/register", publicController.Register)
		public.POST("/uploadFile", controllers.FileUploadController{}.UploadFile)
	}
	return public
}
