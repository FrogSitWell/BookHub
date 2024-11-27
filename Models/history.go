package Models

import "time"

type History struct {
    ID           uint      `gorm:"primaryKey"`
    UserID       uint      `gorm:"not null"`
    User         User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
    BookID       uint      `gorm:"not null"`
    Book         Book      `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE"`
    LastChapter  int       `gorm:"not null"`
    LastReadTime time.Time `gorm:"autoCreateTime"`
}
