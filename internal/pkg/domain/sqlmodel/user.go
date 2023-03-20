package sqlmodel

import (
	"time"
)

type User struct {
	ID             int       `json:"id" gorm:"primarykey"`
	FullName       string    `json:"full_name"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	PasswordHash   string    `json:"-"`
	AvatarFileName string    `json:"avatar_file_name"`
	IsDeleted      bool      `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (m *User) TableName() string {
	return "users"
}
