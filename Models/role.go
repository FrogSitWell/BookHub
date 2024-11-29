package Models

// Role mô hình cho bảng roles
type Role struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"size:255;not null" json:"name"`
	
}
