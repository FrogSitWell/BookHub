package Controllers

import (
	"bookhub/Models"
	"bookhub/database"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)
func GetBooksName(c *gin.Context) {
	// Lấy bookId từ path parameter
    bookID := c.Param("id")
    if bookID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Book ID is required"})
        return
    }

    // Convert bookId thành uint
    bookIDUint, err := strconv.ParseUint(bookID, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Book ID"})
        return
    }

    // Truy vấn database để lấy tên truyện
    var book Models.Book
    if err := database.DB.Select("name").Where("id = ?", uint(bookIDUint)).First(&book).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Book not found", "error": err.Error()})
        return
    }

    // Trả về tên truyện
    c.JSON(http.StatusOK, gin.H{"name": book.Name})
}