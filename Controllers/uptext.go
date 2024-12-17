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
	"gorm.io/gorm"
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
		SortOrder   int64  `json:"sortOrder"`      // sortOrder
	}

	// Parse dữ liệu JSON từ frontend
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error(), "details": request})
		return
	}

	// Kiểm tra BookID tồn tại trong bảng Book
	var book Models.Book
	if err := database.DB.First(&book, request.BookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found", "error": err.Error()})
		return
	}
// Kiểm tra vị trí sort_order có trống không
isAvailable, err := CheckSortOrderAvailable(database.DB, request.BookID, request.SortOrder)
if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"message": "Error checking sort_order", "error": err.Error()})
    return
}

// Nếu sort_order không khả dụng, trả về gợi ý giá trị mới
if !isAvailable {
    // Lấy sort_order lớn nhất hiện tại (MAX(sort_order)) + 1
    var maxSortOrder int64
    err := database.DB.Model(&Models.Chapter{}).
        Where("book_id = ?", request.BookID).
        Select("MAX(sort_order)").Row().Scan(&maxSortOrder)
    if err != nil && err != gorm.ErrRecordNotFound {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching max sort_order", "error": err.Error()})
        return
    }

    suggestedSortOrder := maxSortOrder + 1

    c.JSON(http.StatusBadRequest, gin.H{
        "message":            fmt.Sprintf("Vị trí sort_order %d đã tồn tại", request.SortOrder),
        "suggestedSortOrder": suggestedSortOrder,
    })
    return
}


	// Upload nội dung chương lên Cloudinary và lấy URL
	contentURL, err := UploadToCloudinary(request.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to upload content", "error": err.Error()})
		return
	}

	// Lấy số thứ tự (SortOrder) của chương tiếp theo nếu không có sortOrder
	if request.SortOrder == 0 {
		var lastChapter Models.Chapter
		if err := database.DB.Where("book_id = ?", request.BookID).Order("sort_order desc").First(&lastChapter).Error; err != nil {
			// Nếu không có chương nào, mặc định SortOrder = 1
			request.SortOrder = 1
		} else {
			// Nếu có chương rồi, lấy sort_order tiếp theo
			request.SortOrder = lastChapter.SortOrder + 1
		}
		
	}

	// Tạo chương mới
	newChapter := Models.Chapter{
		BookID:       request.BookID,
		Title:        request.ChapterName,
		ContentURL:   contentURL,
		SortOrder:    request.SortOrder, // Gán SortOrder theo giá trị được chọn
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

// Kiểm tra xem vị trí sort_order có trống hay không
func CheckSortOrderAvailable(db *gorm.DB, bookID uint, sortOrder int64) (bool, error) {
	var chapter Models.Chapter
	// Kiểm tra xem sort_order đã có trong sách chưa
	if err := db.Where("book_id = ? AND sort_order = ?", bookID, sortOrder).First(&chapter).Error; err != nil {
		// Nếu không tìm thấy chương với sort_order này, vị trí có thể sử dụng
		if err == gorm.ErrRecordNotFound {
			return true, nil
		}
		return false, err
	}
	return false, nil
}
