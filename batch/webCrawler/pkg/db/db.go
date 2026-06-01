package db

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() *sqlx.DB {
    dsn := os.Getenv("DB_DSN")
    db  := sqlx.MustConnect("postgres", dsn)
    return db
}
