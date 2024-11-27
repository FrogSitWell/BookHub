package Models

import "time"

type Comment struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      `gorm:"not null"`
    User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
    BookID    uint      `gorm:"not null"`
    Book      Book      `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE"`
    Content   string    `gorm:"type:text;not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
}
