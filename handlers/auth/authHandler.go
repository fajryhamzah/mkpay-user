package auth

import (
	"errors"
	"fmt"

	"github.com/fajryhamzah/mkpay-user/handlers/auth/store"
	"github.com/fajryhamzah/mkpay-user/utils/cache"
	"github.com/fajryhamzah/mkpay-user/utils/jwt"
	"golang.org/x/crypto/bcrypt"

	"github.com/fajryhamzah/mkpay-user/src/user"
)

//AuthHandler handler
type AuthHandler struct {
	userRepo user.RepoInterface
}

//Auth handler
func (auth AuthHandler) Auth(email string, password string) (interface{}, error) {
	user := auth.userRepo.FindByEmail(email)
	code := user.GetCode()
	userType := user.GetUserType()

	if code == "" || !user.GetActive() {
		return nil, errors.New("User not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(password))

	if nil != err {
		return nil, errors.New("Wrong password")
	}

	return generateTokenResponse(code, userType)
}

func (auth AuthHandler) GetNewTokenWithRefreshToken(token string) (interface{}, error) {
	cacheClient := cache.Get()

	savedToken, err := cacheClient.Read(token)

	if nil != err {
		return nil, fmt.Errorf("%s", "Invalid refresh token")
	}

	err = cacheClient.Delete(token)

	if nil != err {
		return nil, fmt.Errorf("%s", "Try again later")
	}

	payload, err := jwt.ParseToken(savedToken)

	if nil != err {
		return nil, err
	}

	user := auth.userRepo.FindByCode(payload.GetCode())

	if "" == user.GetCode() {
		return nil, fmt.Errorf("%s", "User not found")
	}

	return generateTokenResponse(user.GetCode(), user.GetUserType())
}

func generateTokenResponse(code string, userType string) (interface{}, error) {
	token, err := jwt.EncodeAuth(code, userType, store.GetTokenLifeTime())

	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.EncodeAuth(code, userType, store.GetRefreshTokenLifeTime())

	if err != nil {
		return nil, err
	}

	err = store.SaveRefreshToken(refreshToken)

	if err != nil {
		return nil, err
	}

	return authHandlerTransformer{
		token,
		refreshToken,
	}, nil
}

//NewAuthHandler init handler
func NewAuthHandler(userRepo user.RepoInterface) AuthHandler {
	return AuthHandler{userRepo}
}
