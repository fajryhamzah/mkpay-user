package controllers

import (
	"net/http"

	stringHelper "github.com/fajryhamzah/mkpay-user/utils/helpers/string"
	tokenHelper "github.com/fajryhamzah/mkpay-user/utils/helpers/token"
	"github.com/julienschmidt/httprouter"
)

func (m MainController) AddUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		internalError(w)
		return
	}

	email := r.PostFormValue("email")
	pass := r.PostFormValue("password")
	userType := r.PostFormValue("user_type")
	phoneNumber := r.PostFormValue("phone_number")

	if !validateAddUserRequest(w, email, pass, userType, phoneNumber) {
		return
	}

	result, err := m.userHandler.AddUser(email, pass, userType, phoneNumber)

	if nil != err {
		requestNotValid(w, []string{err.Error()})
		return
	}

	successResponse(w, result)
}

func (m MainController) ActivateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		internalError(w)
		return
	}

	payload := tokenHelper.GetPayloadToken(r)
	userCode := r.PostFormValue("code")

	if !validateActiveRequest(w, payload.GetCode(), userCode) {
		return
	}

	err := m.userHandler.ActivateUser(userCode, true)

	if nil != err {
		requestNotValid(w, []string{err.Error()})
		return
	}

	successResponse(w, "ok")
	return
}

func (m MainController) DeactivateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		internalError(w)
		return
	}

	payload := tokenHelper.GetPayloadToken(r)
	userCode := r.PostFormValue("code")

	if !validateActiveRequest(w, payload.GetCode(), userCode) {
		return
	}

	err := m.userHandler.ActivateUser(userCode, false)

	if nil != err {
		requestNotValid(w, []string{err.Error()})
		return
	}

	successResponse(w, "ok")
	return
}

func (m MainController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		internalError(w)
		return
	}

	payload := tokenHelper.GetPayloadToken(r)
	userCode := p.ByName("code")

	if !validateActiveRequest(w, payload.GetCode(), userCode) {
		return
	}

	err := m.userHandler.DeleteUser(userCode)

	if nil != err {
		requestNotValid(w, []string{err.Error()})
		return
	}

	successResponse(w, "ok")
	return
}

func validateActiveRequest(w http.ResponseWriter, code string, userCode string) bool {
	if code == userCode {
		requestNotValid(w, []string{"Can't do this action to yourself"})
		return false
	}

	return true
}

func validateAddUserRequest(w http.ResponseWriter, email string, password string, userType string, phoneNumber string) bool {
	var msg []string

	if !stringHelper.IsEmailValid(email) {
		msg = append(msg, "Email is not valid")
	}

	if len(password) < 7 {
		msg = append(msg, "Password must have 7 characters minimum")
	}

	if !stringHelper.IsValidUserType(userType) {
		msg = append(msg, "Invalid user type")
	}

	if len(phoneNumber) < 11 {
		msg = append(msg, "Phone number not valid")
	}

	if len(msg) > 0 {
		requestNotValid(w, msg)

		return false
	}

	return true
}
