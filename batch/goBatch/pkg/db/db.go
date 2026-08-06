package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() *sqlx.DB {
	dbHost     := os.Getenv("DB_HOST")
	dbPort     := os.Getenv("DB_PORT")
	dbUser     := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName 	   := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
	)
    db, err := sqlx.Connect("postgres", dsn)

	if err == nil {
		return db
	} else {
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_SERVICE_NAME"),
			dbPort,
			dbUser,
			dbPassword,
			dbName,
		)
		db = sqlx.MustConnect("postgres", dsn)
	}
    return db
}
