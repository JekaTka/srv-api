package config

import "os"

type DB struct {
	Host     string
	User     string
	Database string
	Password string
}

type Config struct {
	DB *DB
}

func Load() *Config {
	return &Config{
		DB: &DB{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Database: os.Getenv("DB_NAME"),
			Password: os.Getenv("DB_PASSWORD"),
		},
	}
}
