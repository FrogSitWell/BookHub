package Controllers

import (
	"bookhub/Models"
	"bookhub/database"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Lấy thông tin chi tiết sách
func Getinfo(c *gin.Context) {
	bookID := c.Param("id")
    var book Models.Book

    // Tìm sách theo ID và preload cả thông tin Genre và Status
    if err := database.DB.Preload("Genre").Preload("Status").First(&book, bookID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "Name": book.Name,
        "Author": book.Author,
        "AvatarURL": book.AvatarURL,
        "Genre": gin.H{
            "Name": book.Genre.Name,
        },
        "TotalChapters": book.TotalChapters,
        "Status": gin.H{
            "StatusName": book.Status.StatusName,
        },
        "AverageRate": book.AverageRate,
    })
}
func GetBooksByAuthor(c *gin.Context) {
    bookID := c.Param("id") // Lấy Book ID từ URL
    var book Models.Book

    // Bước 1: Tìm sách theo Book ID
    if err := database.DB.First(&book, bookID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    authorName := book.Author // Lấy tên tác giả từ sách đã tìm được
    var books []Models.Book

    // Bước 2: Truy vấn danh sách các truyện của tác giả
    if err := database.DB.Where("author = ?", authorName).Find(&books).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Books not found for this author"})
        return
    }

    // Bước 3: Trả về danh sách truyện
    c.JSON(http.StatusOK, gin.H{
        "author": authorName,
        "books":  books,
    })
}
func GetBookDescriptionByID(c *gin.Context) {
    bookID := c.Param("id") // Lấy ID từ URL parameter
    var book Models.Book

    // Tìm sách theo ID
    if err := database.DB.Select("description").First(&book, bookID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    // Trả về description
    c.JSON(http.StatusOK, gin.H{"description": book.Description})
}

func GetTopChapters(c *gin.Context) {
    // Lấy BookID từ path parameter
    bookID := c.Param("id")
    if bookID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Book ID is required"})
        return
    }

    // Chuyển BookID thành uint
    bookIDUint, err := strconv.ParseUint(bookID, 10, 0) // Chuyển thành uint
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Book ID"})
        return
    }

    // Lấy cuốn sách theo BookID và preload các chapters, giới hạn 3 chương có sort_order lớn nhất
    var book Models.Book
    if err := database.DB.Preload("Chapters", func(db *gorm.DB) *gorm.DB {
        return db.Where("book_id = ?", bookIDUint).
            Order("sort_order DESC"). // Sắp xếp theo sort_order giảm dần
            Limit(3)                   // Giới hạn số lượng chương là 3
    }).First(&book, uint(bookIDUint)).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Book not found", "error": err.Error()})
        return
    }

    // Trả về danh sách chương đã được sắp xếp
    c.JSON(http.StatusOK, gin.H{
        "chapters": book.Chapters,
    })
}


func GetAllChapters(c *gin.Context) {
    // Lấy BookID từ path parameter
    bookID := c.Param("id")
    if bookID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Book ID is required"})
        return
    }

    // Chuyển BookID thành uint
    bookIDUint, err := strconv.ParseUint(bookID, 10, 0) // Chuyển thành uint
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Book ID"})
        return
    }

    // Lấy cuốn sách theo BookID và preload các chapters, sắp xếp theo sort_order
    var book Models.Book
    if err := database.DB.Preload("Chapters", func(db *gorm.DB) *gorm.DB {
        return db.Order("sort_order ASC") // Sắp xếp theo sort_order
    }).First(&book, uint(bookIDUint)).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Book not found", "error": err.Error()})
        return
    }

    // Trả về danh sách chương đã được sắp xếp
    c.JSON(http.StatusOK, gin.H{
        "chapters": book.Chapters,
    })
}

