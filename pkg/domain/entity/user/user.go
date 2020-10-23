package user

import (
	"github.com/JekaTka/srv-api/pkg/domain/entity"
)

type User struct {
	entity.Model
	Nickname  string
	Email     string
	Password  string
}
