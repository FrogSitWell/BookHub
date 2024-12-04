package Controllers

import (
	"bookhub/Models"
	"bookhub/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Lấy thông tin người dùng (Profile)
func GetProfile(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	var user Models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
