package psql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Driver struct {
	Connection string
}

func (p Driver) GetConnection() *sql.DB {
	db, err := sql.Open("postgres", p.Connection)

	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Opening postgresql connection.....")
	return db
}
