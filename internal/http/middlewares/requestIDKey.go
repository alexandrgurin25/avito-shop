package middlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)


const (
	RequestIDKey string = "request_id"
)

// Мидлварь на генерацию уникального request_id
func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()

		ctx := context.WithValue(r.Context(), RequestIDKey, requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
