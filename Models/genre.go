package Models

import "time"

type Genre struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"type:varchar(255);unique;not null"`
    Books     []Book    `gorm:"foreignKey:GenreID"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
