package middlewares

import (
	"net/http"
	"os"
)

func Auth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do something before
		apiToken := os.Getenv("API_TOKEN")
		token := r.Header.Get("token")

		if apiToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("Unauthorized"))
			return
		}

		if token != apiToken {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("Unauthorized"))
			return
		}

		handler.ServeHTTP(w, r)
	})
}
