package routes

import (
	"bookhub/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/register", Controllers.RegisterUser)
	}
}