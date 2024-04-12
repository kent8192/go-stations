package middleware

import (
	"context"
	"net/http"
)

type ContextKey string

func Device(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		k := ContextKey("User-Agent")
		ctx := context.WithValue(context.Background(), k, r.UserAgent())
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
