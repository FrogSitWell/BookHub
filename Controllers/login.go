package Controllers

import (
	"bookhub/Models"
	"bookhub/config"
	"bookhub/database"
	"errors"
	"strconv"
	"time"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


func GenerateJWT(user Models.User) (string, error) {
	claims := &jwt.StandardClaims{
		Subject:   strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JwtKey)
}
// Hàm đăng nhập người dùng
func LoginUser(c *gin.Context) {
	var user Models.User

	// Lấy dữ liệu từ request body
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format"})
		return
	}

	// Kiểm tra xem người dùng có tồn tại trong cơ sở dữ liệu hay không
	var existingUser Models.User
	if err := database.DB.Where("name = ?", user.Name).First(&existingUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respondWithError(c, http.StatusUnauthorized, "Invalid name or password")
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error checking user"})
		return
	}

	// So sánh mật khẩu người dùng nhập với mật khẩu đã hash trong cơ sở dữ liệu
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
		respondWithError(c, http.StatusUnauthorized, "Invalid  password")
		return
	}

	// Tạo JWT token
	token, err := generateJWT(existingUser)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Error generating token")
		return
	}

	// Phản hồi với thông tin người dùng và token
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    existingUser,
		"token":   token,
	})
}
