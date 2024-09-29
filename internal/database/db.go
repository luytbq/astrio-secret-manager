package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/luytbq/astrio-secret-manager/config"
)

func NewPosgresDB() (*sql.DB, error) {
	// connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.App.PG_USER,
		config.App.PG_PASSWORD,
		config.App.PG_HOST,
		config.App.PG_PORT,
		config.App.PG_DB_NAME,
		config.App.PG_SSL_MODE,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = initPostgres(db)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func initPostgres(db *sql.DB) error {
	err := db.Ping()

	if err != nil {
		return err
	}

	log.Println("Postgres DB init successfully")

	return nil
}
