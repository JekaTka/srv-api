package user

import "github.com/JekaTka/cryptohex-api/pkg/domain/entity"

//Reader interface
type Reader interface {
	//Get(id entity.ID) (*User, error)
	//SearchByEmail(email string) (*User, error)
	//List() ([]*User, error)
}

//Writer user writer
type Writer interface {
	Create(e *User) (entity.ID, error)
	//Update(e *User) error
	//Delete(id entity.ID) error
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase use case interface
type Service interface {
	Reader
	Writer
}
