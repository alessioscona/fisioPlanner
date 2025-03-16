package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := "postgres://fisioplanner:fisioplanner@localhost/fisioplanner?sslmode=disable"
	return sql.Open("postgres", connStr)
}
