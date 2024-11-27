package Controllers

import (
	"bookhub/Models"
	"bookhub/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Đăng ký người dùng mới
func RegisterUser(c *gin.Context) {
	var user Models.User

	// Bind dữ liệu từ request vào struct User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid data",
		})
		return
	}

	// Kiểm tra xem email đã tồn tại chưa
	var existingUser Models.User
	if err := database.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Email already registered",
		})
		return
	}

	// Hash mật khẩu trước khi lưu
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error hashing password",
		})
		return
	}

	// Lưu người dùng vào cơ sở dữ liệu
	user.Password = string(hashedPassword) // Lưu mật khẩu đã băm
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating user",
		})
		return
	}

	// Trả về phản hồi thành công
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}
