package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() *sqlx.DB {
    dsn := os.Getenv("DB_DSN")
log.Printf("DB_DSN=[%s]", dsn)
    db  := sqlx.MustConnect("postgres", dsn)
    return db
}
