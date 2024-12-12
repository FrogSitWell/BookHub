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
		api.DELETE("/:id", Controllers.DeleteBook)
		api.POST("/chapters", Controllers.AddChapter) 
		api.GET("/name/:id",Controllers.GetBooksName)
		api.GET("/info/:id",Controllers.GetIaN)
		api.GET("/chapters/:id",Controllers.GetChapters)
		api.DELETE("/:id/chapters/:chapterID", Controllers.DeleteChapter)
		api.GET("/genres", Controllers.GetGenres)       // Lấy danh sách thể loại
		api.DELETE("/genres/:id", Controllers.DeleteGenre)
		api.PUT("/genres/:id", Controllers.UpdateGenre)
		api.POST("/genres", Controllers.CreateGenre)
		api.POST("/banners",Controllers.UploadBanner)
		api.GET("/banners",Controllers.GetBanners)
		api.DELETE("/banners/:id",Controllers.DeleteBanner)
		api.GET("/users", Controllers.GetUsers)
		api.PUT("/users/:id", Controllers.UpdateUser)
		api.DELETE("/users/:id", Controllers.DeleteUser)
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
		bookRoutes.DELETE("/:id", Controllers.DeleteBook)
		bookRoutes.PUT("/:id",Controllers.UpdateBook)
		
	}
	// Route để hiển thị trang chỉnh sửa sách
	r.GET("/auth/fixtruyen/:id", Controllers.RenderFixBookPage)
	r.GET("/auth/newchap/:id", Controllers.RenderFixBookPage)
	r.GET("/auth/listchap/:id", Controllers.RenderFixBookPage)
}
