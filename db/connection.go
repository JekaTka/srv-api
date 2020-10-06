package db

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/JekaTka/cryptohex-api/config"

	// blank import to have driver registered
	_ "github.com/lib/pq"
)

func NewConnection(cfg *config.DB) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Database)
	log.Println("DB conn string: ", connStr)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	log.Println("DB CONNECTED!")

	return conn, nil
}
