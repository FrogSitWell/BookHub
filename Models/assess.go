package Models

import "time"

type Assess struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      `gorm:"not null"`
    User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
    BookID    uint      `gorm:"not null"`
    Book      Book      `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE"`
    Rating    int       `gorm:"check:rating >= 1 AND rating <= 5;not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
}
