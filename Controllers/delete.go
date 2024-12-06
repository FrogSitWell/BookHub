package Controllers

import (
	"bookhub/Models"
	"bookhub/database"

	"net/http"

	"github.com/gin-gonic/gin"
)
func DeleteBook(c *gin.Context) {
    bookID := c.Param("id") // Lấy ID từ URL

    var book Models.Book
    if err := database.DB.First(&book, bookID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy sách"})
        return
    }

    if err :=  database.DB.Delete(&book).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Xóa sách thất bại"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Xóa sách thành công"})
}
