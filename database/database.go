package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB là biến toàn cục để sử dụng trong các nơi khác
var DB *gorm.DB

func Connect() {
	// Chuỗi kết nối (DSN - Data Source Name)
	dsn := "root:123456789@tcp(127.0.0.1:3306)/bookhub?charset=utf8mb4&parseTime=True&loc=Local"

	// Kết nối với MySQL bằng GORM
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Không thể kết nối tới cơ sở dữ liệu:", err)
	}

	fmt.Println("Kết nối thành công tới cơ sở dữ liệu!")
}
