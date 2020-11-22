package server

import (
	"fmt"
	"net/http"

	"github.com/fajryhamzah/mkpay-user/controllers"
	"github.com/fajryhamzah/mkpay-user/controllers/middlewares"
)

//Server web server struct
type Server struct {
	controller controllers.MainController
	middleware middlewares.AuthMiddleware
}

//InitServer initialize web server
func (s Server) InitServer() {
	fmt.Println("Running webserver.....")
	http.ListenAndServe(":8081", s.getRoute())
}

func New(
	c controllers.MainController,
	m middlewares.AuthMiddleware,
) Server {
	return Server{c, m}
}
