package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Auth handler
func (m MainController) Auth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	if err := r.ParseForm(); err != nil {
		internalError(w)
		return
	}

	email := r.PostFormValue("email")
	pass := r.PostFormValue("password")

	if !authRequestValidator(w, email, pass) {
		return
	}

	fmt.Fprintf(w, "LOGIN!")
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
