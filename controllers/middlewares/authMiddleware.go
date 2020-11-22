package middlewares

import (
	"net/http"
	"strings"

	"github.com/fajryhamzah/mkpay-user/controllers"
	"github.com/fajryhamzah/mkpay-user/handlers/auth/store"
	"github.com/julienschmidt/httprouter"
)

func (a AuthMiddleware) ValidateToken(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if err := r.ParseForm(); err != nil {
			controllers.InternalError(w)
			return
		}

		authHeader := r.Header.Get("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			msg := []string{"Authorization empty"}
			controllers.RequestNotValid(w, msg)
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)

		if !store.IsTokenValidAndRoleValid(tokenString, a.userRepo) {
			msg := []string{"Invalid Request"}
			controllers.RequestNotValid(w, msg)
			return
		}

		next(w, r, ps)
	}
}

func (a AuthMiddleware) ContentMustBeJson(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r, ps)
	}
}
