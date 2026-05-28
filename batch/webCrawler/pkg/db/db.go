package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
    godotenv.Load()
}

func Connect() *sqlx.DB {
    dsn := os.Getenv("DB_DSN")

    db, err := sqlx.Open("postgres", dsn)
    if err != nil {
        log.Fatal("DB接続失敗:", err)
    }

    if err := db.Ping(); err != nil {
        log.Fatal("DB疎通失敗:", err)
    }

    return db
}
