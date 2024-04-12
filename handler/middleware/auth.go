package middleware

import (
	"net/http"
	"os"
)

func Auth(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		targetID := os.Getenv("ID")
		targetPass := os.Getenv("PASSWORD")
		id, pass, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "An error occurred during authentication", http.StatusBadRequest)
			return
		}
		if id != targetID {
			http.Error(w, "ID unmatched", http.StatusUnauthorized)
			return
		}
		if pass != targetPass {
			http.Error(w, "Password unmatched", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
