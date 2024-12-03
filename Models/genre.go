package Models

type Genre struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255);unique;not null"`
}
