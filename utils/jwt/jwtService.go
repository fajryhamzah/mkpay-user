package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	jwtgo "github.com/dgrijalva/jwt-go"
)

type claims struct {
	Code string `json:"code"`
	Role string `json:"role"`
	jwt.StandardClaims
}

func EncodeAuth(code string, role string, lifeTime int64) (string, error) {
	duration, _ := time.ParseDuration(fmt.Sprintf("%ds", lifeTime))
	claims := claims{
		code,
		role,
		jwtgo.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}

	token := jwtgo.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_AUTH")))
}

func ParseToken(token string) (*JwtPayload, error) {
	tokenParse, err := jwtgo.Parse(token, func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return []byte(os.Getenv("JWT_AUTH")), nil
	})

	if nil != err {
		return nil, fmt.Errorf("%s", "Invalid Token")
	}

	claims, ok := tokenParse.Claims.(jwtgo.MapClaims)

	if !ok || !tokenParse.Valid {
		return nil, fmt.Errorf("%s", "Invalid Token")
	}

	userCode := claims["code"].(string)
	role := claims["role"].(string)

	return &JwtPayload{userCode, role}, nil
}
