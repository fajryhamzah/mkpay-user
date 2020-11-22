package middlewares

import (
	"net/http"
	"strings"

	"github.com/fajryhamzah/mkpay-user/controllers"
	"github.com/fajryhamzah/mkpay-user/src/user"
	"github.com/fajryhamzah/mkpay-user/utils/helpers/token"
	"github.com/julienschmidt/httprouter"
)

func (a AuthMiddleware) AdminOnly(next httprouter.Handle) httprouter.Handle {
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

		payload := token.GetPayloadToken(r)

		if payload.GetRole() != user.RoleAdmin {
			controllers.Unauthorized(w)
			return
		}

		next(w, r, ps)
	}
}
