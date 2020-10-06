package config

import (
	"time"
	"os"
)

type DB struct {
	Host     string
	User     string
	Database string
	Password string
}

type Server struct {
	Port int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

type Config struct {
	DB *DB
	Server *Server
}

func Load() *Config {
	return &Config{
		DB: &DB{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Database: os.Getenv("DB_NAME"),
			Password: os.Getenv("DB_PASSWORD"),
		},
		Server: &Server{
			Port: 1323,
			ReadTimeout: 10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}
