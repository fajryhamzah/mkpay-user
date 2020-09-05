package user

import (
	"testing"
)

const dummyID = 1
const dummyName = "Test Name"
const dummyPassword = "Password"
const dummyUserType = TypeAdmin
const dummyActive = true
const dummyUserCode = "ae93e1b6-6f70-4d09-b184-4381245e74d6"

var user User = User{
	ID:       dummyID,
	Email:    dummyName,
	UserType: dummyUserType,
	Active:   dummyActive,
	Code:     dummyUserCode,
}

func TestGetID(t *testing.T) {
	if user.GetID() != dummyID {
		t.Error("Wrong User Id")
	}
}

func TestGetEmail(t *testing.T) {
	if user.GetEmail() != dummyName {
		t.Error("Wrong User Email")
	}
}

func TestSetPassword(t *testing.T) {
	user.SetPassword(dummyPassword)
	if user.GetPassword() != dummyPassword {
		t.Errorf("Wrong User Password, should be %s got %s instead", user.GetPassword(), dummyPassword)
	}
}

func TestGetUserType(t *testing.T) {
	if user.GetUserType() != dummyUserType {
		t.Error("Wrong User type")
	}
}

func TestGetCode(t *testing.T) {
	if user.GetCode() != dummyUserCode {
		t.Error("Wrong User code")
	}
}

func TestGetActive(t *testing.T) {
	if user.GetActive() != dummyActive {
		t.Error("Wrong User active")
	}
}
