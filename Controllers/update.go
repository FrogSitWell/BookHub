package Controllers

import (
	"bookhub/Models"
	"bookhub/database"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)
func RenderFixBookPage(c *gin.Context) {
    bookID := c.Param("id") // Sử dụng Param thay vì Query nếu lấy từ đường dẫn
    if bookID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID sách không hợp lệ"})
        return
    }
    log.Printf("Đang truy cập ID: %s", bookID)
  
    // Lấy sách từ cơ sở dữ liệu
    var book Models.Book
    if err := database.DB.Where("id = ?", bookID).First(&book).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy sách"})
        return
    }
   
    // Render template với dữ liệu sách
    c.HTML(http.StatusOK, "fixtruyen.html", gin.H{
        "Book": book,
        "Genres": []Models.Genre{ /* Load danh sách thể loại từ DB */ },
    })
}
