package routes

import (
	"bookhub/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	
	
	api := r.Group("/api")
	
	{
		api.POST("/register", Controllers.RegisterUser)
		api.POST("/login", Controllers.LoginUser)
		api.GET("/genre",Controllers.GetGenre)
	}
	

	// Route bảo vệ (cần có token)
	protected := r.Group("/protected")
	protected.Use(Controllers.TokenAuthMiddleware())
	{
		// Các route cần bảo vệ
		protected.GET("/profile", Controllers.GetProfile)
		protected.GET("/booklist", Controllers.GetBooksByUser)
		protected.POST("/upload",Controllers.CreateBookWithAvatar)
	}
	// Routes cho sách
	bookRoutes := api.Group("/books")
	{
		bookRoutes.POST("/", Controllers.CreateBook)
		bookRoutes.GET("/", Controllers.GetBooks)
		bookRoutes.GET("/:id", Controllers.GetBookByID)
		bookRoutes.PUT("/:id", Controllers.UpdateBook)
		bookRoutes.DELETE("/:id", Controllers.DeleteBook)
	

	}

}
