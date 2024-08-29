package middleware

import (
	"context"
	"mezink/stdlib/log"
	"net/http"

	"github.com/google/uuid"
)

func LoggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("x-request-id")
		if id == "" {
			id = uuid.NewString()
		}
		ctx := context.WithValue(r.Context(), "x-request-id", id)
		r = r.WithContext(ctx)

		log.InfoContext(r.Context(), "Request [%s] %q", r.Method, r.URL.String())
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
