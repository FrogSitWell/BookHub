package Controllers

import (
	"bookhub/Models"
	"bookhub/database"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)
func GetBooksWithChapters(c *gin.Context) {
    var books []Models.Book

    // Lấy danh sách truyện từ dưới lên với điều kiện 50 chương trở lên
    result := database.DB.Preload("Genre").Where("total_chapters >= ?", 50).Order("id DESC").Find(&books)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    // Trả danh sách truyện dưới dạng JSON
    c.JSON(http.StatusOK, books)
}

// Tạo hàm tính khoảng thời gian
func timeAgo(t time.Time) string {
	duration := time.Since(t)
	switch {
	case duration < time.Minute:
		return "Vừa xong"
	case duration < time.Hour:
		return fmt.Sprintf("%d phút trước", int(duration.Minutes()))
	case duration < 24*time.Hour:
		return fmt.Sprintf("%d giờ trước", int(duration.Hours()))
	case duration < 7*24*time.Hour:
		return fmt.Sprintf("%d ngày trước", int(duration.Hours()/24))
	default:
		return t.Format("02/01/2006")
	}
}

// API Lấy 10 chương mới
func GetLatestChapters(c *gin.Context) {
	var chapters []Models.Chapter

	// Lấy 10 chương mới nhất và join với bảng Book và Genre
	if err := database.DB.
		Preload("Book").
		Preload("Book.Genre").
		Order("updated_at DESC").
		Limit(10).
		Find(&chapters).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Chuẩn bị dữ liệu phản hồi
	var response []map[string]interface{}
	for _, chapter := range chapters {
		response = append(response, map[string]interface{}{
			"chapter_id":    chapter.ID,
			"chapter_title": chapter.Title,
			"chapter_number": chapter.ChapterNumber,
			"book_name":     chapter.Book.Name,
			"author":        chapter.Book.Author,
			"genre":         chapter.Book.Genre.Name,
			"time_ago":      timeAgo(chapter.UpdatedAt),
		})
	}

	c.JSON(http.StatusOK, response)
}

