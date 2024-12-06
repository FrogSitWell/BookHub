package Controllers

import (
	"bookhub/Models"
	"bookhub/database"

	"net/http"

	"github.com/gin-gonic/gin"
)

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
	bookID := c.Param("id")
	var book Models.Book

	// Tìm sách theo ID
	if err := database.DB.Preload("Genre").First(&book, bookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// Cập nhật thông tin sách
// Xử lý yêu cầu PUT để cập nhật sách
func UpdateBook(c *gin.Context) {
    var updatedBook Models.Book
    bookId := c.Param("id")

    if err := c.ShouldBindJSON(&updatedBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Dữ liệu không hợp lệ"})
        return
    }

    // Kiểm tra xem sách có tồn tại không
    var book Models.Book
    if result := database.DB.First(&book, bookId); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Sách không tồn tại"})
        return
    }

    // Cập nhật sách
    book.Name = updatedBook.Name
    book.Author = updatedBook.Author
    book.Description = updatedBook.Description
    book.GenreID = updatedBook.GenreID
    book.AvatarURL = updatedBook.AvatarURL

    if result := database.DB.Save(&book); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Có lỗi khi cập nhật sách"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Cập nhật thành công"})
}

//theloai
func GetGenre(c *gin.Context) {
    var genres []Models.Genre
    if err := database.DB.Model(&Models.Genre{}).Find(&genres).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot fetch genres"})
        return
    }
    c.JSON(http.StatusOK, genres) // Trả về danh sách thể loại cùng với sách
}