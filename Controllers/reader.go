package Controllers

import (
	"bookhub/Models"
	"bookhub/database"

	"net/http"

	"github.com/gin-gonic/gin"
)

// SearchBooks tìm kiếm sách theo từ khóa trong tên sách và tác giả
func SearchBooks(c *gin.Context) {
    query := c.DefaultQuery("query", "") // Lấy giá trị từ query string

    if query == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Query is required"})
        return
    }

    var books []Models.Book
    // Tìm kiếm sách theo tên sách và tác giả
	database.DB.Where("name LIKE ? OR author LIKE ?", "%"+query+"%", "%"+query+"%").
        Preload("Genre").Preload("User").Find(&books)

    // Trả về kết quả tìm kiếm
    c.JSON(http.StatusOK, gin.H{
        "books": books,
    })
}