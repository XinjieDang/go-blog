package main

import (
	"dxj/configs"
	_ "dxj/docs" //重要一定要引入对应的docs，不然swagger会报错
	"dxj/internal/initialize"
	"dxj/internal/router"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           blog API
// @version         1.0
// @description     go 博客系统
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:7001
// @BasePath  /api/v1
// @securityDefinitions.basic  BasicAuth
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	initialize.GlobalInit()
	r := router.Init()
	// init swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(configs.EnvConfig.Server.Port)
}
