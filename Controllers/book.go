package Controllers

import (
	"bookhub/Models"
	"bookhub/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Thêm sách mới
func CreateBook(c *gin.Context) {
	var book Models.Book

	// Parse JSON input
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Lưu sách vào cơ sở dữ liệu
	if err := database.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully", "data": book})
}

// Lấy danh sách sách
func GetBooks(c *gin.Context) {
	var books []Models.Book

	// Lấy tất cả sách từ cơ sở dữ liệu
	if err := database.DB.Preload("Genre").Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

	c.JSON(http.StatusOK, books)
}

// Lấy thông tin chi tiết sách
func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var book Models.Book

	// Tìm sách theo ID
	if err := database.DB.Preload("Genre").First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// Cập nhật thông tin sách
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book Models.Book

	// Kiểm tra sách có tồn tại
	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Parse JSON input
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cập nhật thông tin sách
	if err := database.DB.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully", "data": book})
}

// Xóa sách
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book Models.Book

	// Kiểm tra sách có tồn tại
	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Xóa sách
	if err := database.DB.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
