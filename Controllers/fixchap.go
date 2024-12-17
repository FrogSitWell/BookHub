package Controllers

import (
	"bookhub/Models"
	"bookhub/database"
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetChapterByID(c *gin.Context) {
    // Lấy ID chương từ URL
    chapterId := c.Param("chapterId")  // Sửa thành đúng tên tham số trong URL
    id, err := strconv.Atoi(chapterId)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter ID"})
        return
    }

    // Tìm chương từ database
    var chapter Models.Chapter
    if err := database.DB.First(&chapter, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Chapter not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        }
        return
    }

    // Trả về dữ liệu chương
    c.JSON(http.StatusOK, chapter)
}
func UpdateToCloudinary(content string) (string, error) {
	// Khởi tạo Cloudinary client
	cld, err := cloudinary.NewFromURL("cloudinary://548828148668285:UP4RJ7PESLfae1utwV0-Xhn3Oww@dwttzicez")
	if err != nil {
		return "", fmt.Errorf("cloudinary setup error: %v", err)
	}

	// Tạo context
	ctx := context.Background()

	// Upload nội dung trực tiếp từ chuỗi
	uploadResult, err := cld.Upload.Upload(ctx, strings.NewReader(content), uploader.UploadParams{
		Folder:       "chapters",                                // Lưu vào thư mục "chapters"
		ResourceType: "raw",                                    // Định dạng dữ liệu là "raw"
		PublicID:     fmt.Sprintf("chapter-%d", time.Now().Unix()), // Tên tệp duy nhất dựa trên timestamp
	})
	if err != nil {
		return "", fmt.Errorf("error uploading content to Cloudinary: %v", err)
	}

	// In ra kết quả trả về từ Cloudinary (nếu cần kiểm tra)
	fmt.Println("Cloudinary Response:", uploadResult)

	// Trả về URL đã lưu trữ trên Cloudinary
	return uploadResult.URL, nil
}

func UpdateChapter(c *gin.Context) {
    // Lấy ID chương từ URL
    chapterId := c.Param("chapterId")
    id, err := strconv.Atoi(chapterId)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter ID"})
        return
    }

    // Tìm chương trong database để đảm bảo chapterID là hợp lệ và đồng bộ
    var chapter Models.Chapter
    if err := database.DB.First(&chapter, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Chapter not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        }
        return
    }

    // Ràng buộc dữ liệu từ request body
    var input Models.Chapter
    if err := c.ShouldBindJSON(&input); err != nil {
        // In ra chi tiết lỗi để dễ dàng debug
        fmt.Println("Error binding JSON:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
        return
    }

    // Kiểm tra nếu có nội dung chương mới, tải lên Cloudinary và lấy URL
    var contentURL string
    if input.ContentURL != "" {
        contentURL, err = UpdateToCloudinary(input.ContentURL) // Upload nội dung lên Cloudinary
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to upload content", "error": err.Error()})
            return
        }
    }

    // Cập nhật dữ liệu chương
    chapter.Title = input.Title
    chapter.SortOrder = input.SortOrder
    chapter.ContentURL = contentURL // Cập nhật URL mới

    // Lưu vào cơ sở dữ liệu
    if err := database.DB.Save(&chapter).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update chapter"})
        return
    }

    // Trả về dữ liệu sau khi cập nhật
    c.JSON(http.StatusOK, chapter)
}
