package db

import "database/sql"

//PSQL driver postgresql
const PSQL = "psql"

//DbInterface database driver interface
type DbInterface interface {
	GetConnection() *sql.DB
	Close()
}
