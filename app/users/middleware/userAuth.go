package middleware

import (
	"context"
	"net/http"
	"strings"
)

// Прослойка для проверки подлиноости пользователя
func UserAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenType := "Basic "
		token := strings.Split(r.Header.Get("Authorization"), tokenType)
		if len(token) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Empty token\n"))
			return
		}
		ctx := context.WithValue(r.Context(), "Token", token[1])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
