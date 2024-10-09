package web

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Error   bool   `json:"error"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func SucessResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	response := Response{
		Error:   false,
		Data:    data,
		Message: "",
	}
	json.NewEncoder(w).Encode(&response)
}

func ErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	response := Response{
		Error:   true,
		Data:    nil,
		Message: message,
	}
	json.NewEncoder(w).Encode(&response)
}

func GenericResponse(w http.ResponseWriter, code int, data any, message string) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(code)

	err := false
	if code > 399 {
		err = true
	}

	response := Response{
		Error:   err,
		Data:    data,
		Message: message,
	}
	json.NewEncoder(w).Encode(&response)
}
