package controllers

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	ErrorCode int      `json:"code"`
	Msg       []string `json:"message"`
}

func response(w http.ResponseWriter, code int, message interface{}) {
	if code != 200 {
		w.WriteHeader(code)
	}

	json.NewEncoder(w).Encode(message)
}

func errResponse(w http.ResponseWriter, err errorResponse) {
	response(w, err.ErrorCode, err)
}

func internalError(w http.ResponseWriter) {
	response(w, 500, errorResponse{500, []string{"Something happen!"}})
}

func requestNotValid(w http.ResponseWriter, msg []string) {
	response(w, 400, errorResponse{400, msg})
}

func successResponse(w http.ResponseWriter, result interface{}) {
	response(w, 200, result)
}

func notFoundResponse(w http.ResponseWriter, msg []string) {
	response(w, 404, errorResponse{404, msg})
}

func InternalError(w http.ResponseWriter) {
	internalError(w)
}

func RequestNotValid(w http.ResponseWriter, msg []string) {
	requestNotValid(w, msg)
}

func Unauthorized(w http.ResponseWriter) {
	response(w, 403, errorResponse{403, []string{"Unauthorized"}})
}
