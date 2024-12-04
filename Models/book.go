package Models

import "time"

type Book struct {
    ID            uint        `gorm:"primaryKey"`
    Name          string      `gorm:"type:varchar(255);not null"`
    Author        string      `gorm:"type:varchar(255);not null"`
    AvatarURL     string      `gorm:"type:text;not null"`
    Description   string      `gorm:"type:text;not null"`
    TotalChapters int         `gorm:"not null"`
    GenreID       uint        `gorm:"not null"`
    Genre         Genre       `gorm:"foreignKey:GenreID;constraint:OnDelete:CASCADE"`
    UserID        uint        `gorm:"not null"`
    AverageRate   float64     `gorm:"default:0"`
    User          User        `gorm:"foreignKey:UserID"`
    Comments      []Comment   `gorm:"foreignKey:BookID"`
    Assess        []Assess    `gorm:"foreignKey:BookID"`
    Follows       []Follow    `gorm:"foreignKey:BookID"`
    Chapters      []Chapter   `gorm:"foreignKey:BookID"`
    History       []History   `gorm:"foreignKey:BookID"`
    CreatedAt     time.Time   `gorm:"autoCreateTime"`
    UpdatedAt     time.Time   `gorm:"autoUpdateTime"`
}
