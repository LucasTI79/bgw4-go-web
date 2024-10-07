package web

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Message string       `json:"message"`
	Data    *interface{} `json:"data"`
	Error   bool         `json:"error"`
}

func ResponseJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response := &response{}
	response.Data = &data

	if code > 399 {
		response.Error = true
	}

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
