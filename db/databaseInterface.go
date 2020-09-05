package db

import "database/sql"

const PSQL = "psql"

type DbInterface interface {
	GetConnection() *sql.DB
}
