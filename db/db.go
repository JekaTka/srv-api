package db

import (
	"fmt"
	"log"

	"github.com/JekaTka/cryptohex-api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	*gorm.DB
}

func NewConnection(cfg *config.DB) (*DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Database)
	log.Println("DB conn string: ", connStr)

	conn, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	log.Println("DB CONNECTED!")

	return &DB{conn}, nil
}
