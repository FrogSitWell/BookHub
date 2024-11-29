package Models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"size:255;not null" json:"name"`
	Password string `gorm:"size:255;not null" json:"password"`
	Email    string `gorm:"size:255;unique;not null" json:"email"`
	RoleID   uint   `gorm:"not null" json:"role_id"`       // Khóa ngoại đến bảng Role
	Role     Role   `gorm:"foreignKey:RoleID" json:"role"` // Liên kết với bảng Role
}
