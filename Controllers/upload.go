package Controllers

import (
	"bookhub/Models"
	"bookhub/database"
	"net/http"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)
func CreateBookWithAvatar(c *gin.Context) {
	// 1. Lấy thông tin từ request
	name := c.PostForm("name")
	description := c.PostForm("description")
	author := c.PostForm("author")
	genreID, err := strconv.Atoi(c.PostForm("genreid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid genre ID"})
		return
	}

	// Lấy userID từ context (được gán trong middleware TokenAuthMiddleware)
	userIDRaw := c.MustGet("userID")
	userID, ok := userIDRaw.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID in context"})
		return
	}

	// Chuyển đổi userID từ string sang uint
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Kiểm tra xem userID có tồn tại trong bảng users không
	var user Models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Lấy status_id từ form, mặc định là 1 nếu không được cung cấp
	statusID := c.PostForm("statusid")
	if statusID == "" {
		statusID = "1" // Mặc định là 1
	}
	statusIDInt, err := strconv.Atoi(statusID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status ID"})
		return
	}

	// Lấy file từ form
	file, _, err := c.Request.FormFile("avatar")
	if err != nil && err != http.ErrMissingFile { // Không bắt buộc phải có ảnh
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get image", "details": err.Error()})
		return
	}

	// 2. Lưu thông tin sách vào cơ sở dữ liệu
	book := Models.Book{
		Name:         name,
		Description:  description,
		Author:       author,
		GenreID:      uint(genreID),
		StatusID:     uint(statusIDInt),
		UserID:       uint(userIDInt),
		TotalChapters: 0, // Mặc định là 0
	}
	if err := database.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	// 3. Nếu có file ảnh, upload lên Cloudinary và cập nhật URL
	if file != nil {
		cld, err := cloudinary.NewFromURL("cloudinary://548828148668285:UP4RJ7PESLfae1utwV0-Xhn3Oww@dwttzicez")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize Cloudinary"})
			return
		}

		resp, err := cld.Upload.Upload(c, file, uploader.UploadParams{Folder: "books"})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image", "details": err.Error()})
			return
		}

		// Cập nhật URL ảnh vào sách
		book.AvatarURL = resp.SecureURL
		if err := database.DB.Save(&book).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book with image"})
			return
		}
	}

	// 4. Trả về thông tin sách đã tạo
	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully", "data": book})
}

