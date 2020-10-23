package user

import (
	"github.com/JekaTka/srv-api/pkg/services/password"
	"time"
	"github.com/JekaTka/srv-api/pkg/domain/entity"
)

//Service service interface
type service struct {
	repo Repository
	pwd  password.Passwrder
}

//NewService create new use case
func NewService(r Repository, pwd password.Passwrder) Service {
	return &service{
		repo: r,
		pwd:  pwd,
	}
}

//Create an user
func (s *service) Create(e *User) (entity.ID, error) {
	now := time.Now()
	e.ID = entity.NewID()
	e.CreatedAt = now
	e.UpdatedAt = now
	pwd, err := s.pwd.Generate(e.Password)
	if err != nil {
		return e.ID, err
	}
	e.Password = pwd

	return s.repo.Create(e)
}
