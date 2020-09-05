package server

import (
	"fmt"
	"net/http"
)

func getRoute() {
	http.HandleFunc("/", HelloServer)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
