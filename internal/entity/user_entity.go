package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	FullName        string      `gorm:"size:255;not null" json:"name"`
	Username    string         `gorm:"size:50; not null" json:"username"`
	Password    string         `gorm:"size:255;not null" json:"password"`
	BirthDate   time.Time      `gorm:"type:date;" json:"birthDate"`
	Email       string         `gorm:"size:255;uniqueIndex;not null" json:"email"`
	PhoneNumber string         `gorm:"size:20;" json:"phoneNumber"`
	Address     string         `gorm:"size:512;" json:"address"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
