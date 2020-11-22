package user

import "database/sql"

//TypeAdmin user type
const TypeAdmin = "admin"

//TypeStudent user type
const TypeStudent = "student"

const RoleAdmin = "admin"
const RoleUser = "user"

//ModelInterface user model interface
type ModelInterface interface {
	GetID() uint32
	GetEmail() string
	GetPassword() string
	GetUserType() string
	GetActive() bool
	GetCode() string
	GetPhoneNumber() string
	SetPassword(password string)
	SetActive(active bool)
	New(row *sql.Row)
}
