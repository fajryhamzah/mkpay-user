package token

import (
	"net/http"
	"strings"

	"github.com/fajryhamzah/mkpay-user/utils/jwt"
)

func GetPayloadToken(r *http.Request) *jwt.JwtPayload {
	authHeader := r.Header.Get("Authorization")
	tokenString := strings.Replace(authHeader, "Bearer ", "", -1)

	payload, _ := jwt.ParseToken(tokenString)

	return payload
}
