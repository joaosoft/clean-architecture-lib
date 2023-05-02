package postgres

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	_ "github.com/lib/pq"              // postgres driver
)

func NewConnection(driver, dataSource string) (db *sql.DB, err error) {
	return sql.Open(driver, dataSource)
}
