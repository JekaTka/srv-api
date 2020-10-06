package user

import (
	"database/sql"
	"github.com/JekaTka/cryptohex-api/pkg/domain/entity"
)

//PGRepo repo
type PGRepo struct {
	db *sql.DB
}

//NewPGRepoRepository create new repository
func NewPGRepoRepository(db *sql.DB) *PGRepo {
	return &PGRepo{
		db: db,
	}
}

//Create an user
func (r *PGRepo) Create(e *User) (entity.ID, error) {
	stmt, err := r.db.Prepare(`
		INSERT INTO USERS (id, nickname, email, password, created_at, updated_at) 
		VALUES(?,?,?,?,?,?)`)
	if err != nil {
		return e.ID, err
	}

	_, err = stmt.Exec(
		e.ID,
		e.Nickname,
		e.Email,
		e.Password,
		e.CreatedAt,
		e.UpdatedAt,
	)
	if err != nil {
		return e.ID, err
	}

	err = stmt.Close()
	if err != nil {
		return e.ID, err
	}

	return e.ID, nil
}
