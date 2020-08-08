package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	*gorm.DB
}

func NewConnection() (*DB, error) {
	// Get database details from environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, DBName)
	log.Println("DB conn string: ", connStr)

	conn, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	log.Println("DB CONNECTED!")

	return &DB{conn}, nil
}
