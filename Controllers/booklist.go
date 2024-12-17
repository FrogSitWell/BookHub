package Controllers

import (
	"bookhub/Models"
	"bookhub/database"
	"net/http"

	"github.com/gin-gonic/gin"
)
func GetBooksByUser(c *gin.Context) {
    // Lấy userID từ context
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusBadRequest, gin.H{"message": "User ID not found"})
        return
    }

    // Chuyển sang kiểu string nếu cần
    userIDStr, ok := userID.(string)
    if !ok {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid User ID"})
        return
    }

    var books []Models.Book

    // Truy vấn sách với Preload Genre và Status
    if err := database.DB.Preload("Genre").Preload("Status").Where("user_id = ?", userIDStr).Find(&books).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve books"})
        return
    }

    // Nếu không có sách nào
    if len(books) == 0 {
        c.JSON(http.StatusOK, gin.H{"message": "No books found"})
        return
    }

    // Trả về danh sách sách
    c.JSON(http.StatusOK, gin.H{"books": books})
}
