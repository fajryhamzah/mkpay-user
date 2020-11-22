package user

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

//Repository user
type Repository struct {
	db *sql.DB
}

func (r Repository) findBy(where string, value interface{}) ModelInterface {
	var result ModelInterface = &User{}

	row := r.db.QueryRow(fmt.Sprintf("SELECT * FROM \"user\" WHERE %s = %s AND deleted_at is null LIMIT 1", where, value))

	result.New(row)

	return result
}

//FindByCode get user model by user code
func (r Repository) FindByCode(code string) ModelInterface {
	return r.findBy("code", fmt.Sprintf("'%s'", code))
}

//FindByEmail get user model by email
func (r Repository) FindByEmail(email string) ModelInterface {
	return r.findBy("email", fmt.Sprintf("'%s'", email))
}

func (r Repository) Save(model ModelInterface) error {
	sqlStatement := `INSERT INTO "user" (code, email, password, phone_number, user_type, active) VALUES ($1, $2, $3, $4, $5, $6)`
	code := uuid.New().String()
	_, err := r.db.Exec(sqlStatement, code, model.GetEmail(), model.GetPassword(), model.GetPhoneNumber(), model.GetUserType(), model.GetActive())

	return err
}

func (r Repository) Update(id uint32, model ModelInterface) error {
	sqlStatement := `UPDATE "user" SET email=$1, password=$2, phone_number=$3, user_type=$4, active=$5 WHERE id=$6`

	_, err := r.db.Exec(sqlStatement, model.GetEmail(), model.GetPassword(), model.GetPhoneNumber(), model.GetUserType(), model.GetActive(), id)

	return err
}

func (r Repository) Delete(id uint32) error {
	sqlStatement := `DELETE FROM "user" WHERE id=$1`

	_, err := r.db.Exec(sqlStatement, id)

	return err
}

//New init
func New(db *sql.DB) RepoInterface {
	return Repository{db}
}
