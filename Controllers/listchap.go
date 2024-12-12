package Controllers

import (
	"bookhub/Models"
	"bookhub/database"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func GetIaN(c *gin.Context) {
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

    // Truy vấn database để lấy tên truyện và URL ảnh
	var book Models.Book
	if err := database.DB.Select("name, avatar_url").Where("id = ?", uint(bookIDUint)).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found", "error": err.Error()})
		return
	}

	// Trả về tên truyện và URL ảnh
	c.JSON(http.StatusOK, gin.H{
		"name":       book.Name,
		"avatar_url": book.AvatarURL,
	})
}
func GetChapters(c *gin.Context) {
    // Lấy BookID từ path parameter
    bookID := c.Param("id")
    if bookID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Book ID is required"})
        return
    }

    // Chuyển BookID thành uint
    bookIDUint, err := strconv.ParseUint(bookID, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Book ID"})
        return
    }

    // Lấy cuốn sách theo BookID
    var book Models.Book
    if err := database.DB.Preload("Chapters").First(&book, uint(bookIDUint)).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Book not found", "error": err.Error()})
        return
    }

    // Trả về danh sách chương
    c.JSON(http.StatusOK, gin.H{
        "chapters": book.Chapters,
    })
}
func DeleteChapter(c *gin.Context) {
    // Lấy bookID và chapterID từ URL
    bookID := c.Param("id")       // Lấy bookID từ tham số "id"
    chapterID := c.Param("chapterID") // Lấy chapterID từ tham số "chapterID"
    var book Models.Book
    // Kiểm tra xem bookID và chapterID có hợp lệ không
    if bookID == "" || chapterID == "" {
        c.JSON(400, gin.H{"error": "Thiếu thông tin ID sách hoặc chương"})
        return
    }

    // Tìm chương cần xóa
    var chapter Models.Chapter
    if err := database.DB.Where("id = ? AND book_id = ?", chapterID, bookID).First(&chapter).Error; err != nil {
        c.JSON(404, gin.H{"error": "Chương không tồn tại trong sách"})
        return
    }

    // Xóa chương
    if err := database.DB.Where("id = ? AND book_id = ?", chapterID, bookID).Delete(&Models.Chapter{}).Error; err != nil {
        c.JSON(500, gin.H{"error": "Không thể xóa chương"})
        return
    }

    // Cập nhật thứ tự các chương còn lại
    if err := database.DB.Model(&Models.Chapter{}).
        Where("book_id = ? AND sort_order > ?", bookID, chapter.SortOrder).
        Update("sort_order", gorm.Expr("sort_order - ?", 1)).Error; err != nil {
        c.JSON(500, gin.H{"error": "Không thể cập nhật thứ tự chương"})
        return
    }
    if err := database.DB.Delete(&chapter).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete chapter", "error": err.Error()})
        return
    }
    
    // Cập nhật lại tổng số chương trong sách
    book.TotalChapters -= 1
    if book.TotalChapters < 0 {
        book.TotalChapters = 0 // Đảm bảo tổng số chương không nhỏ hơn 0
    }
    
    if err := database.DB.Save(&book).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update book's total chapters", "error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"message": "Xóa chương thành công"})
}
