package Controllers

import (
	"bookhub/Models"
	"bookhub/database"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Lấy danh sách sách
func GetBooklists(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 10  // Mỗi trang chứa 10 sách
	offset := (page - 1) * limit

	var books []Models.Book
	var totalBooks int64

	// Lấy tổng số sách để tính tổng trang
	if err := database.DB.Model(&Models.Book{}).Count(&totalBooks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count total books"})
		return
	}

	// Lấy sách theo offset và limit
	if err := database.DB.Preload("Genre").Offset(offset).Limit(limit).Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"books": books,
		"total": totalBooks,  // Trả về tổng số sách để tính toán số trang
	})
}

func FilterBooks(c *gin.Context) {
	var books []Models.Book
	query := database.DB.Preload("Genre").Preload("Status")

	// Lọc theo Status
    if statusID := c.Query("status_id"); statusID != "" {
        query = query.Where("status_id = ?", statusID)
    }

    // Lọc theo Genre
    if genreID := c.Query("genre_id"); genreID != "" {
        if id, err := strconv.Atoi(genreID); err == nil {
            query = query.Where("genre_id = ?", id)
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid genre_id"})
            return
        }
    }

    // Lọc theo số chương lớn hơn giá trị được chỉ định
    if minChapters := c.Query("min_chapters"); minChapters != "" {
        if min, err := strconv.Atoi(minChapters); err == nil {
            query = query.Where("total_chapters >= ?", min)
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid min_chapters value"})
            return
        }
    }

	// Thực thi query
	if err := query.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

	c.JSON(http.StatusOK, books)
}


func GetFilterOptions(c *gin.Context) {
    // Lấy các giá trị status và genre từ cơ sở dữ liệu
    var status []Models.Status
    var genre []Models.Genre
    
    // Lấy các status và genres
    if err := database.DB.Find(&status).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch statuses"})
        return
    }
    
    if err := database.DB.Find(&genre).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch genres"})
        return
    }
    
    // Trả về dữ liệu status và genre cho frontend
    c.JSON(http.StatusOK, gin.H{
        "statuses": status,
        "genres":   genre,
    })
}

