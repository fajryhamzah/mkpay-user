package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (m MainController) Auth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		internalError(w)
		return
	}

	email := r.PostFormValue("email")
	pass := r.PostFormValue("password")

	if !authRequestValidator(w, email, pass) {
		return
	}

	result, err := m.authHandler.Auth(email, pass)

	if err != nil {
		notFoundResponse(w, []string{err.Error()})
		return
	}

	successResponse(w, result)
}

func (m MainController) TokenByRefreshToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		internalError(w)
		return
	}

	refreshToken := r.PostFormValue("refresh_token")

	result, err := m.authHandler.GetNewTokenWithRefreshToken(refreshToken)

	if err != nil {
		requestNotValid(w, []string{err.Error()})
		return
	}

	successResponse(w, result)
}

func authRequestValidator(w http.ResponseWriter, email string, pass string) bool {
	var msg []string

	if email == "" {
		msg = append(msg, "Email must be filled.")
	}

	if pass == "" {
		msg = append(msg, "Password must be filled.")
	}

	if !isEmailValid(email) {
		msg = append(msg, "Wrong email format.")
	}

	if len(msg) > 0 {
		requestNotValid(w, msg)

		return false
	}

	return true
}
