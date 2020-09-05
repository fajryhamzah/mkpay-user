package user

//TypeAdmin user type
const TypeAdmin = "admin"

//TypeStudent user type
const TypeStudent = "student"

//ModelInterface user model interface
type ModelInterface interface {
	GetID() uint32
	GetEmail() string
	GetPassword() string
	GetUserType() string
	GetActive() bool
	GetCode() string
	SetPassword(password string)
}
