package main

import (
	"bookhub/Models"
	"bookhub/database"
	"bookhub/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)
func main() {
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
	// Khởi tạo Gin router
	r := gin.Default()

	// Cấu hình các routes
	routes.SetupRoutes(r)

	// Chạy Gin server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Không thể chạy server: %v", err)
	}
}