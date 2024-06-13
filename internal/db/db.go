package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	PublicHost  string
	Port        string
	DBUser      string
	DBPassword  string
	DBAddress   string
	DBName      string
	JWTSecret   string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err!= nil{
		log.Fatal("error loading .env file")
	}
	return Config{
		PublicHost: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBAddress:   fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		DBName:      os.Getenv("DB_NAME"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
	}
}

func NewSQLStorage(cfg Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBAddress, cfg.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

