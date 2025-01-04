package Models

import "time"

type History struct {
    ID           uint      `gorm:"primaryKey"`
    UserID       uint      `gorm:"not null"`
    User         User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
    BookID       uint      `gorm:"not null"`
    Book         Book      `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE"`
    ChapterID    uint      `gorm:"not null"` // Thay thế LastChapter
    Chapter      Chapter   `gorm:"foreignKey:ChapterID;constraint:OnDelete:CASCADE"`
    LastReadTime time.Time `gorm:"autoCreateTime"`
}
