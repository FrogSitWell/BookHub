package Controllers

import (
	"bookhub/Models"
	"bookhub/database"
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)
func UploadToCloudinary(content string) (string, error) {
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

func AddChapter(c *gin.Context) {
	var request struct {
		BookID      uint   `json:"bookId"`         // bookId
		ChapterName string `json:"chapterName"`    // chapterName
		Content     string `json:"content"`        // content
	}

	// Parse dữ liệu JSON từ frontend
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error(), "details": request})
		return
	}

	fmt.Println("Request nhận được từ frontend:", request)

	// Kiểm tra BookID tồn tại trong bảng Book
	var book Models.Book
	if err := database.DB.First(&book, request.BookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found", "error": err.Error()})
		return
	}

	// Upload nội dung chương lên Cloudinary và lấy URL
	contentURL, err := UploadToCloudinary(request.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to upload content", "error": err.Error()})
		return
	}

	// In ra URL nội dung đã được upload lên Cloudinary
	fmt.Println("Content URL from Cloudinary:", contentURL)

	// Lấy số chương tiếp theo (ChapterNumber)
	var lastChapter Models.Chapter
	if err := database.DB.Where("book_id = ?", request.BookID).Order("chapter_number desc").First(&lastChapter).Error; err != nil {
		// Nếu không có chương nào, mặc định ChapterNumber = 1
		lastChapter.ChapterNumber = 0
	}

	// Tạo chương mới
	newChapter := Models.Chapter{
		BookID:       request.BookID,
		Title:        request.ChapterName,
		ContentURL:   contentURL,
		ChapterNumber: lastChapter.ChapterNumber + 1, // Tăng số chương
	}

	// Thêm chương mới vào cơ sở dữ liệu
	if err := database.DB.Create(&newChapter).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create chapter", "error": err.Error()})
		return
	}

	// Cập nhật tổng số chương trong sách
	book.TotalChapters += 1
	if err := database.DB.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update book's total chapters", "error": err.Error()})
		return
	}

	// Trả về thành công
	c.JSON(http.StatusOK, gin.H{"success": true, "chapter": newChapter})
}
