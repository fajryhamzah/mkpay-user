package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s Server) getRoute() *httprouter.Router {
	controller := s.controller
	router := httprouter.New()

	router.GET("/:name", HelloServer)
	router.POST("/auth", controller.Auth)

	return router
}

//HelloServer dummy homepage
func HelloServer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello, %s!", ps.ByName("name"))
}
