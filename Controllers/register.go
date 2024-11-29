package Controllers

import (
	"bookhub/Models"
	"bookhub/database"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Chìa khóa bí mật để ký JWT token (nên thay đổi trong môi trường sản xuất)
var jwtKey = []byte("secret")

// Hàm tiện ích để phản hồi lỗi API
func respondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"message": message})
}

// Đăng ký người dùng mới
func RegisterUser(c *gin.Context) {
	var user Models.User

	// Lấy dữ liệu từ request
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON format",
		})
		return
	}

	// Gán vai trò mặc định nếu RoleID là 0
    if user.RoleID == 0 {
        user.RoleID = 3 // Gán vai trò mặc định là Reader (ID = 3)
    }


	log.Printf("Received User: %+v\n", user) // Log giá trị user nhận được

	// In ra thông tin user đã nhận được
    fmt.Println("Received User:", user)

	// Kiểm tra role_id có tồn tại trong bảng roles không
	var role Models.Role
	if err := database.DB.First(&role, user.RoleID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid role_id: role does not exist",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error checking role_id",
		})
		return
	}
	// In ra thông tin role đã nhận được
    fmt.Println("Received Role:", role)
	// Kiểm tra xem email đã tồn tại chưa
	var existingUser Models.User
	if err := database.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		respondWithError(c, http.StatusConflict, "Email already registered")
		return
	}

	// Kiểm tra độ dài mật khẩu
	if len(user.Password) < 6 {
		respondWithError(c, http.StatusBadRequest, "Password must be at least 6 characters long")
		return
	}

	// Hash mật khẩu trước khi lưu
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Error hashing password")
		return
	}
	user.Password = string(hashedPassword)

	// Lưu người dùng mới vào cơ sở dữ liệu
	if err := database.DB.Create(&user).Error; err != nil {
		respondWithError(c, http.StatusInternalServerError, "Error creating user")
		return
	}

	// Tạo JWT token
	token, err := generateJWT(user)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Error generating token")
		return
	}

	// Phản hồi thành công
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"user":    user,
		"token":   token,
	})
}

// Hàm tạo JWT token
func generateJWT(user Models.User) (string, error) {
	claims := &jwt.StandardClaims{
		Subject:   strconv.Itoa(int(user.ID)),                 // ID người dùng
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),      // Token hết hạn sau 1 ngày
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
