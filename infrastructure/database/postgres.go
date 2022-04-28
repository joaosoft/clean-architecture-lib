package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	_ "github.com/lib/pq"              // postgres driver
)

func NewDatabase() (db *sql.DB, err error) {
	return sql.Open("postgres", "postgres://foursource:password@localhost:5432?sslmode=disable")
}
