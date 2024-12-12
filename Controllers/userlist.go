package Controllers

import (
	"bookhub/Models"
	"bookhub/database"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Lấy danh sách người dùng
func GetUsers(c *gin.Context) {
	var users []Models.User
	if err := database.DB.Preload("Role").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
// Sửa thông tin người dùng
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user Models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Bind dữ liệu từ request body vào user
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cập nhật thông tin người dùng
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
// Xóa người dùng
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&Models.User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
