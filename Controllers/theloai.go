package Controllers

import (
	"bookhub/Models"
	"bookhub/database"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Lấy danh sách thể loại
func GetGenres(c *gin.Context) {
	var genres []Models.Genre
	if err := database.DB.Find(&genres).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lấy danh sách thể loại"})
		return
	}
	c.JSON(http.StatusOK, genres)
}
func DeleteGenre(c *gin.Context) {
	id := c.Param("id")

	// Xóa thể loại
	if err := database.DB.Delete(&Models.Genre{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể xóa thể loại"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Xóa thành công"})
}
func UpdateGenre(c *gin.Context) {
	id := c.Param("id")
	var genre Models.Genre

	// Kiểm tra thể loại có tồn tại không
	if err := database.DB.First(&genre, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Thể loại không tồn tại"})
		return
	}

	// Gán dữ liệu mới
	if err := c.ShouldBindJSON(&genre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}

	// Lưu vào database
	if err := database.DB.Save(&genre).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể cập nhật thể loại"})
		return
	}

	c.JSON(http.StatusOK, genre)
}
func CreateGenre(c *gin.Context) {
    var newGenre Models.Genre

    // Ràng buộc dữ liệu từ request JSON vào struct
    if err := c.ShouldBindJSON(&newGenre); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
        return
    }

    // Kiểm tra tên thể loại có bị trùng không
    if err := database.DB.Where("name = ?", newGenre.Name).First(&Models.Genre{}).Error; err == nil {
        c.JSON(http.StatusConflict, gin.H{"error": "Thể loại đã tồn tại"})
        return
    }

    // Thêm thể loại vào database
    if err := database.DB.Create(&newGenre).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo thể loại"})
        return
    }

    // Phản hồi thành công
    c.JSON(http.StatusCreated, newGenre)
}
