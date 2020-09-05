package user

const TypeAdmin = "admin"
const TypeStudent = "student"

type Interface interface {
	GetID() uint32
	GetEmail() string
	GetPassword() string
	GetUserType() string
	GetActive() bool
	GetCode() string
	SetPassword(password string)
}
