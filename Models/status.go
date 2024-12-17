package Models

type Status struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	StatusName string `gorm:"type:varchar(50);not null"` // Ví dụ: Đang ra, Tạm dừng, Hoàn thành

}