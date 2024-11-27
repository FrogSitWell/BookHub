package Models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:255;not null"`
	Password string `gorm:"size:255;not null"`
	Email    string `gorm:"size:255;unique;not null"`
	RoleID   uint   `gorm:"not null"`          // Khóa ngoại đến bảng Role
	Role     Role   `gorm:"foreignKey:RoleID"` // Liên kết với bảng Role
}
