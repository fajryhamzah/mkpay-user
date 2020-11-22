package user

import (
	"errors"

	"github.com/fajryhamzah/mkpay-user/src/user"
)

type UserHandler struct {
	userRepo user.RepoInterface
}

func (u UserHandler) AddUser(email string, password string, userType string, phoneNumber string) (interface{}, error) {
	userObject := &user.User{
		Email:       email,
		UserType:    userType,
		PhoneNumber: phoneNumber,
		Active:      true,
	}

	userObject.SetPassword(password)

	err := u.userRepo.Save(userObject)
	if nil != err {
		return nil, err
	}

	return []string{"Successfully registered"}, nil
}

func (u UserHandler) ActivateUser(code string, isActive bool) error {
	user := u.userRepo.FindByCode(code)

	if "" == user.GetCode() {
		return errors.New("User not found")
	}

	user.SetActive(isActive)

	err := u.userRepo.Update(user.GetID(), user)
	if nil != err {
		return err
	}

	return nil
}

func (u UserHandler) DeleteUser(code string) error {
	user := u.userRepo.FindByCode(code)

	if "" == user.GetCode() {
		return errors.New("User not found")
	}

	err := u.userRepo.Delete(user.GetID())

	if nil != err {
		return err
	}

	return nil
}

func NewUserHandler(userRepo user.RepoInterface) UserHandler {
	return UserHandler{userRepo}
}
