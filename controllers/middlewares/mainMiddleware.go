package middlewares

import "github.com/fajryhamzah/mkpay-user/src/user"

type AuthMiddleware struct {
	userRepo user.RepoInterface
}

func NewAuthMiddleware(userRepo user.RepoInterface) AuthMiddleware {
	return AuthMiddleware{userRepo: userRepo}
}
