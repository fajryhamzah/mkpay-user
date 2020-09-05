package user

//User model
type User struct {
	ID       uint32
	Email    string
	UserType string
	Active   bool
	Code     string
	password string
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

//SetPassword password setter
func (u *User) SetPassword(password string) {
	u.password = password
}
