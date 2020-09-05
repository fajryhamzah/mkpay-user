package user

type User struct {
	ID       uint32
	Email    string
	UserType string
	Active   bool
	Code     string
	password string
}

func (u User) GetID() uint32 {
	return u.ID
}

func (u User) GetCode() string {
	return u.Code
}

func (u User) GetEmail() string {
	return u.Email
}

func (u User) GetPassword() string {
	return u.password
}

func (u User) GetActive() bool {
	return u.Active
}

func (u User) GetUserType() string {
	return u.UserType
}

func (u *User) SetPassword(password string) {
	u.password = password
}
