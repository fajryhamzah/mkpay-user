package controllers

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	ErrorCode int      `json:"code"`
	Msg       []string `json:"message"`
}

func errResponse(w http.ResponseWriter, err errorResponse) {
	w.WriteHeader(err.ErrorCode)
	json.NewEncoder(w).Encode(err)
}

func internalError(w http.ResponseWriter) {
	w.WriteHeader(500)
	json.NewEncoder(w).Encode(errorResponse{500, []string{"Something happen!"}})
}

func requestNotValid(w http.ResponseWriter, msg []string) {
	w.WriteHeader(400)
	json.NewEncoder(w).Encode(errorResponse{400, msg})
}
