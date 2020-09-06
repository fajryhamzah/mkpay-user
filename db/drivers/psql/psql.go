package psql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//Driver database driver
type Driver struct {
	Connection string
	db         *sql.DB
}

//GetConnection get postgresql connection
func (p Driver) GetConnection() *sql.DB {
	db, err := sql.Open("postgres", p.Connection)

	if err != nil {
		panic(err)
	}

	p.db = db

	fmt.Println("Opening postgresql connection.....")
	return db
}

//Close database connection
func (p Driver) Close() {
	p.db.Close()
}
