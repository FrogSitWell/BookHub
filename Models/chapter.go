package Models

import "time"

type Chapter struct {
    ID            uint      `gorm:"primaryKey"`
    BookID        uint      `gorm:"not null"`
    Book          Book      `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE"`
    ChapterNumber int       `gorm:"not null"`
    Title         string    `gorm:"type:varchar(255);not null"`
    SortOrder     int64     `gorm:"not null"`
    ContentURL    string    `gorm:"type:text;not null"`
    CreatedAt     time.Time `gorm:"autoCreateTime"`
}
