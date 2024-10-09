package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func CheckTime(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do something before
		fmt.Println("Request received at", time.Now())

		// call handler
		handler.ServeHTTP(w, r)
	})
}
