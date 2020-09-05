package server

import "net/http"

func InitServer() {
	getRoute()
	http.ListenAndServe(":8080", nil)
}
