package user

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//User model
type User struct {
	ID              uint32
	Email           string
	UserType        string
	PhoneNumber     string
	Active          bool
	Code            string
	password        string
	phoneVerifiedAt *time.Time
	deletedAt       *time.Time
}

//GetID id getter
func (u User) GetID() uint32 {
	return u.ID
}

//GetCode code getter
func (u User) GetCode() string {
	return u.Code
}

//GetEmail email getter
func (u User) GetEmail() string {
	return u.Email
}

//GetPassword password getter
func (u User) GetPassword() string {
	return u.password
}

//GetActive user active getter
func (u User) GetActive() bool {
	return u.Active
}

//GetUserType user type getter
func (u User) GetUserType() string {
	return u.UserType
}

func (u User) GetPhoneNumber() string {
	return u.PhoneNumber
}

//SetPassword password setter
func (u *User) SetPassword(password string) {
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	u.password = string(hashedPass)
}

func (u *User) SetActive(active bool) {
	u.Active = active
}

//New initialize all params from query result
func (u *User) New(row *sql.Row) {
	row.Scan(&u.ID, &u.Code, &u.Email, &u.password, &u.PhoneNumber, &u.phoneVerifiedAt, &u.UserType, &u.Active, &u.deletedAt)
}
