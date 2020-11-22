package server

import (
	"github.com/julienschmidt/httprouter"
)

func (s Server) getRoute() *httprouter.Router {
	controller := s.controller
	router := httprouter.New()
	middlewares := s.middleware

	router.POST("/auth", middlewares.ContentMustBeJson(controller.Auth))
	router.POST("/refresh", middlewares.ContentMustBeJson(controller.TokenByRefreshToken))

	//user
	router.PUT("/user", middlewares.ContentMustBeJson(middlewares.ValidateToken(middlewares.AdminOnly(controller.AddUser))))
	router.POST("/user/activate", middlewares.ContentMustBeJson(middlewares.ValidateToken(middlewares.AdminOnly(controller.ActivateUser))))
	router.POST("/user/deactivate", middlewares.ContentMustBeJson(middlewares.ValidateToken(middlewares.AdminOnly(controller.DeactivateUser))))
	router.DELETE("/user/:code", middlewares.ContentMustBeJson(middlewares.ValidateToken(middlewares.AdminOnly(controller.DeleteUser))))

	return router
}
