package main

import (
	"bookhub/Models"
	"bookhub/database"
	"bookhub/routes"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
func main() {

	// Khởi tạo Gin router
	r := gin.Default()
	
	// Middleware cho CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:5500"}, // Thay bằng URL của frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	database.Connect()

	// Thêm log để thông báo bắt đầu migrate
	fmt.Println("Starting migration...")
	
	// Tự động migrate bảng từ các model
	err :=database.DB.AutoMigrate(
        &Models.User{},
		&Models.Role{},
        &Models.Book{},
        &Models.Genre{},
        &Models.Chapter{},
        &Models.Comment{},
        &Models.Assess{},
        &Models.Follow{},
        &Models.History{},
    )
	if err != nil {
        fmt.Println("Migration failed:", err)
    } else {
        fmt.Println("Migration successful!")
    }

	// Cấu hình các routes
	routes.SetupRoutes(r)

	// Chạy Gin server
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Không thể chạy server: %v", err)
	}
	gin.SetMode(gin.ReleaseMode)

}