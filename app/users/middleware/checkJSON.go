package middleware

import (
	"mime"
	"net/http"
)

// Прослойка для проверки Content-Type == application/json
func CheckJSONMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		if contentType != "" || contentType == " " {
			mt, _, err := mime.ParseMediaType(contentType)
			if err != nil {
				http.Error(w, "Content-Type incorrect\n", http.StatusBadRequest)
				return
			}
			if mt != "application/json" {
				http.Error(w, "Content-Type is not application/json\n", http.StatusBadRequest)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}
