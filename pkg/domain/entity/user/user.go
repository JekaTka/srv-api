package user

import "time"

// "github.com/jinzhu/gorm"

type User struct {
	// gorm.Model
	ID        string
	Hash      string
	Nickname  string
	Email     string `gorm:"type:varchar(150);unique_index"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
