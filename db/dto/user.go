package dto

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Hash     string `gorm:"unique;not null"`
	Nickname string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func (User) TableName() string {
	return "users"
}
