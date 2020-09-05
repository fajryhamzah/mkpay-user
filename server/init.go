package server

import (
	"fmt"
	"net/http"

	"github.com/fajryhamzah/mkpay-user/controllers"
)

//Server web server struct
type Server struct {
	controller controllers.MainController
}

//InitServer initialize web server
func (s Server) InitServer() {
	fmt.Println("Running webserver.....")
	http.ListenAndServe(":8080", s.getRoute())
}

func New(c controllers.MainController) Server {
	return Server{c}
}
