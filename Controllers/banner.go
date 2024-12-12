package Controllers

import (
	"bookhub/Models"
	"bookhub/database"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)
func UploadBanner(c *gin.Context) {
	// Lấy file từ request
	file, _, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Tạo kết nối Cloudinary
	cld, err := cloudinary.NewFromURL("cloudinary://548828148668285:UP4RJ7PESLfae1utwV0-Xhn3Oww@dwttzicez")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize Cloudinary"})
		return
	}

	// Upload ảnh lên Cloudinary
	resp, err := cld.Upload.Upload(c, file, uploader.UploadParams{Folder: "books"})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image", "details": err.Error()})
		return
	}

	// Cập nhật URL ảnh vào cơ sở dữ liệu
	// Giả sử bạn đang làm việc với mô hình `Banner` trong cơ sở dữ liệu
	banner := Models.Banner{ImageURL: resp.SecureURL}
	if err := database.DB.Create(&banner).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image URL to database"})
		return
	}

	// Trả về URL ảnh đã upload
	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully", "url": resp.SecureURL})
}
// API để lấy danh sách ảnh từ cơ sở dữ liệu
func GetBanners(c *gin.Context) {
	var banners []Models.Banner
	if err := database.DB.Find(&banners).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch banners"})
		return
	}
	c.JSON(http.StatusOK, banners)
}
// Hàm xóa banner khỏi cơ sở dữ liệu
func DeleteBanner(c *gin.Context) {
	// Lấy ID của banner cần xóa từ URL
	id := c.Param("id")

	// Tìm banner theo ID
	var banner Models.Banner
	if err := database.DB.First(&banner, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Banner not found"})
		return
	}

	// Xóa banner khỏi cơ sở dữ liệu
	if err := database.DB.Delete(&banner).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete banner from database"})
		return
	}

	// Trả về thông báo thành công
	c.JSON(http.StatusOK, gin.H{"message": "Banner deleted successfully"})
}