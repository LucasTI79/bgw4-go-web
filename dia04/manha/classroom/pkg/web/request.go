package web

import (
	"encoding/json"
	"net/http"
)

// ETL

func RequestJSON(r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(&data)
}
