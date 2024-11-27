package Models

// Role mô hình cho bảng roles
type Role struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:255;not null"`
	Users       []User `gorm:"foreignKey:RoleID"` // Quan hệ với User
}
