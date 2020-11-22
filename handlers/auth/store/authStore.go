package store

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/fajryhamzah/mkpay-user/src/user"
	"github.com/fajryhamzah/mkpay-user/utils/cache"
	stringHelper "github.com/fajryhamzah/mkpay-user/utils/helpers/string"
	"github.com/fajryhamzah/mkpay-user/utils/jwt"
)

func SaveRefreshToken(value string) error {
	return store(value, value, GetRefreshTokenLifeTime())
}

func store(key string, value string, lifeTime int64) error {
	cacheClient := cache.Get()
	seconds := fmt.Sprintf("%ds", lifeTime)

	duration, _ := time.ParseDuration(seconds)

	return cacheClient.Save(key, value, duration)
}

func IsTokenValidAndRoleValid(token string, repo user.RepoInterface) bool {
	claims, err := jwt.ParseToken(token)

	if nil != err {
		return false
	}

	userCode := claims.GetCode()
	role := claims.GetRole()

	if !stringHelper.IsValidUserType(role) {
		return false
	}

	user := repo.FindByCode(userCode)
	if user.GetCode() != userCode || !user.GetActive() || role != user.GetUserType() {
		return false
	}

	return true
}

func GetTokenLifeTime() int64 {
	envLifeTime := os.Getenv("TOKEN_LIFETIME")
	lifeTime, err := strconv.ParseInt(envLifeTime, 10, 64)

	if nil != err {
		return 600
	}

	return lifeTime
}

func GetRefreshTokenLifeTime() int64 {
	envLifeTime := os.Getenv("REFRESH_TOKEN_LIFETIME")
	lifeTime, err := strconv.ParseInt(envLifeTime, 10, 64)

	if nil != err {
		return 172800
	}

	return lifeTime
}
