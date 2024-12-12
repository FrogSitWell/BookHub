package Models

type Banner struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	ImageURL string `gorm:"type:text;not null"`
}
