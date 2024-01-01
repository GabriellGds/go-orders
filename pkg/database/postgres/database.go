package postgres

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB_NAME     = "DB_NAME"
	DB_HOST     = "DB_HOST"
	DB_USER     = "DB_USER"
	DB_PASSWORD = "DB_PASSWORD"
	DB_PORT     = "DB_PORT"
)

func Connect() (*sqlx.DB, error) {
	godotenv.Load()
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv(DB_HOST), os.Getenv(DB_PORT), os.Getenv(DB_NAME), os.Getenv(DB_USER), os.Getenv(DB_PASSWORD))

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
