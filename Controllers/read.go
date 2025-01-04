package Controllers

import (
	"bookhub/Models"
	"bookhub/database"
	"strconv"
	"time"

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

func GetNewChapter(c *gin.Context) {
    chapterId := c.Param("chapterId")
    id, err := strconv.Atoi(chapterId)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter ID"})
        return
    }

    // Lấy chương hiện tại
    var currentChapter Models.Chapter
    if err := database.DB.First(&currentChapter, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Chapter not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        }
        return
    }

    // Lấy chương trước
var prevChapter Models.Chapter
prevChapterURL := ""
if err := database.DB.
    Where("book_id = ? AND sort_order < ?", currentChapter.BookID, currentChapter.SortOrder).
    Order("sort_order DESC").
    First(&prevChapter).Error; err == nil {
    prevChapterURL =  strconv.Itoa(int(prevChapter.ID))
}

// Lấy chương sau
var nextChapter Models.Chapter
nextChapterURL := ""
if err := database.DB.
    Where("book_id = ? AND sort_order > ?", currentChapter.BookID, currentChapter.SortOrder).
    Order("sort_order ASC").
    First(&nextChapter).Error; err == nil {
    nextChapterURL =  strconv.Itoa(int(nextChapter.ID))
}

    // Trả dữ liệu
    c.JSON(http.StatusOK, gin.H{
        "chapter": map[string]interface{}{
            "chapter_number": currentChapter.ChapterNumber,
            "title":          currentChapter.Title,
            "content_url":    currentChapter.ContentURL,
        },
        "prev_chapter": prevChapterURL,
        "next_chapter": nextChapterURL,
    })
}

func UpdateHistory(c *gin.Context) {
    var req struct {
        UserID    uint `json:"user_id" binding:"required"`
        BookID    uint `json:"book_id" binding:"required"`
        ChapterID uint `json:"chapter_id" binding:"required"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var history Models.History
    err := database.DB.Where("user_id = ? AND book_id = ?", req.UserID, req.BookID).First(&history).Error

    if err != nil {
        // Nếu chưa có lịch sử, tạo mới
        history = Models.History{
            UserID:       req.UserID,
            BookID:       req.BookID,
            ChapterID:    req.ChapterID,
            LastReadTime: time.Now(),
        }
        if err := database.DB.Create(&history).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    } else {
        // Nếu đã có, cập nhật
        history.ChapterID = req.ChapterID
        history.LastReadTime = time.Now()
        if err := database.DB.Save(&history).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    }

    c.JSON(http.StatusOK, gin.H{"message": "History updated successfully"})
}
