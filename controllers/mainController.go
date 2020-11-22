package controllers

import (
	"github.com/fajryhamzah/mkpay-user/handlers/auth"
	"github.com/fajryhamzah/mkpay-user/handlers/user"
)

//MainController main controller struct
type MainController struct {
	authHandler auth.AuthHandler
	userHandler user.UserHandler
}

//New initialize main controller
func New(
	authHandler auth.AuthHandler,
	userHandler user.UserHandler,
) MainController {
	return MainController{authHandler, userHandler}
}
