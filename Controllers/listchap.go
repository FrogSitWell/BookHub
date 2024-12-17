package Controllers

import (
	"bookhub/Models"
	"bookhub/database"
	"errors"
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
    bookIDUint, err := strconv.ParseUint(bookID, 10, 0) // Chuyển thành uint
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Book ID"})
        return
    }

    // Lấy cuốn sách theo BookID và preload các chapters, sắp xếp theo sort_order
    var book Models.Book
    if err := database.DB.Preload("Chapters", func(db *gorm.DB) *gorm.DB {
        return db.Order("sort_order ASC") // Sắp xếp theo sort_order
    }).First(&book, uint(bookIDUint)).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Book not found", "error": err.Error()})
        return
    }

    // Trả về danh sách chương đã được sắp xếp
    c.JSON(http.StatusOK, gin.H{
        "chapters": book.Chapters,
    })
}




func DeleteChapter(c *gin.Context) {
    // Lấy ID chương từ request
    chapterId, err := strconv.Atoi(c.Param("chapterId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "ID chương không hợp lệ", "error": err.Error()})
        return
    }

    // Tạo transaction để đảm bảo tính toàn vẹn dữ liệu
    tx := database.DB.Begin()

    // Tìm chương cần xóa
    var chapter Models.Chapter
    if err := tx.First(&chapter, chapterId).Error; err != nil {
        tx.Rollback()
        if errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"message": "Không tìm thấy chương"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Lỗi truy vấn chương", "error": err.Error()})
        return
    }

    // Lấy thông tin truyện chứa chương này
    var book Models.Book
    if err := tx.First(&book, chapter.BookID).Error; err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Lỗi truy vấn truyện", "error": err.Error()})
        return
    }

    // Xóa chương
    if err := tx.Delete(&chapter).Error; err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Lỗi khi xóa chương", "error": err.Error()})
        return
    }

    // Cập nhật tổng số chương trong truyện
    if book.TotalChapters > 0 {
        if err := tx.Model(&book).Update("total_chapters", book.TotalChapters-1).Error; err != nil {
            tx.Rollback()
            c.JSON(http.StatusInternalServerError, gin.H{"message": "Lỗi khi cập nhật tổng số chương", "error": err.Error()})
            return
        }
    }

    // Commit transaction sau khi thực hiện thành công
    if err := tx.Commit().Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Lỗi khi xác nhận giao dịch", "error": err.Error()})
        return
    }

    // Trả về kết quả thành công
    c.JSON(http.StatusOK, gin.H{
        "message":          "Xóa chương thành công",
        "deletedSortOrder": chapter.SortOrder,
        "remainingChapters": book.TotalChapters - 1,
    })
}

