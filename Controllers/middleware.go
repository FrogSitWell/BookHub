package Controllers

import (
	"bookhub/config"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Middleware xác thực token
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy token từ header Authorization
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Missing token"})
			c.Abort()
			return
		}

		// Loại bỏ phần "Bearer " nếu có trong token
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token format"})
			c.Abort()
			return
		}

		// Parse và kiểm tra token
		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return config.JwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Lưu thông tin user vào context
		if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
			c.Set("userID", claims.Subject) // Lưu userID vào context
		}

		// Tiếp tục xử lý request
		c.Next()
	}
}

