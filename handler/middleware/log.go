package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Log struct {
	Timestamp time.Time `json:"timestamp"`
	Latency   int64     `json:"latency"`
	OS        string    `json:"os"`
}

func AccessLog(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		timestamp := time.Now()
		h.ServeHTTP(w, r)
		defer func(accessDate time.Time) {
			latency := time.Since(accessDate)
			userAgent, ok := r.Context().Value(ContextKey("User-Agent")).(string)
			if !ok {
				userAgent = "Unknown"
			}
			log := Log{
				Timestamp: timestamp,
				Latency:   latency.Milliseconds(),
				OS:        userAgent,
			}
			logJSON, err := json.Marshal(log)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Println(string(logJSON))
		}(timestamp)
	}
	return http.HandlerFunc(fn)
}
