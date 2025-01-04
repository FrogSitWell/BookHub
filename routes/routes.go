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
		api.GET("/genre",Controllers.GetGenre)// lấy thể loại truyện...
		api.GET("/statuses", Controllers.GetStatus)// lấy danh sách trạng thái 
		api.DELETE("/:id", Controllers.DeleteBook)
		api.POST("/chapters", Controllers.AddChapter) 
		api.GET("/name/:id",Controllers.GetBooksName) // lấy tên sách trong getbookname
		api.GET("/info/:id",Controllers.GetIaN)//lấy thông tin sách 
		api.GET("/chapters/:id",Controllers.GetChapters)//lấy danh sách chương tại listchap
		api.GET("/chapters/details/:chapterId",Controllers.GetChapterByID)// lấy chi tiết chương tại fixchap
		api.PUT("/chapters/details/:chapterId", Controllers. UpdateChapter)// update nội dung chi tiết tại fixchap
		api.DELETE("/books/:id/chapters/:chapterId", Controllers.DeleteChapter)
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
		api.GET("/search", Controllers.SearchBooks)
		api.GET("/booklists", Controllers. GetBooklists)
		api.GET("/filter", Controllers. FilterBooks)
		api.GET("/filteroptions", Controllers.GetFilterOptions)
		api.GET("/decu", Controllers.GetBooksWithChapters)//lay truyen de cu trong indexuser
		api.GET("/news", Controllers.GetLatestChapters)// lấy chương mới được post lên trong indexuser
		api.GET("/book/:id", Controllers.Getinfo)//lấy thông tin truyện.
		api.GET("/author/:id", Controllers.GetBooksByAuthor)//lấy thông tin sách cùng tác giả 
		api.GET("/description/:id", Controllers.GetBookDescriptionByID)//lấy mô tả sách 
		api.GET("/newchapter/:id", Controllers.GetTopChapters)
		api.GET("/allchapter/:id", Controllers.GetAllChapters)
		api.GET("/chap/:chapterId", Controllers.GetNewChapter)
		api.POST("/history", Controllers.UpdateHistory)
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
