package user

import (
	"database/sql"
)

//Repository user
type Repository struct {
	db *sql.DB
}

//FindByCode get user model by user code
func (r Repository) FindByCode(code string) ModelInterface {
	var result ModelInterface = &User{}

	row := r.db.QueryRow("SELECT * FROM \"user\" WHERE deleted_at is null LIMIT 1")

	result.New(row)

	return result
}
