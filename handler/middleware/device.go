package middleware

import (
	"context"
	"net/http"
)

type DeviceKey string

func Device(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		k := DeviceKey("User-Agent")
		ctx := context.WithValue(r.Context(), k, r.UserAgent())
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
