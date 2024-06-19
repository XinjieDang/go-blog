package middlewares

import "github.com/gin-gonic/gin"

func LoadMiddlewares(r *gin.Engine) *gin.Engine {
	// Global middlewares
	r.Use(ErrorHandle())
	r.Use(Cors())
	//r.Use(Jwt())
	return r
}
