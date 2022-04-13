package middleware

import "net/http"

// Прослойка для сессий пользователя
func SessionMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
