package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

// Прослойка для проверки подлиноости пользователя
func UserAuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		TokenType := "Basic "
		Token := strings.Split(r.Header.Get("Authorization"), TokenType)
		if len(Token) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			// Когда приложение будет закончено.
			// Нужно не забыть удалить вывод ошибки "Token error\n"
			w.Write([]byte("Token error\n"))
			return
		}
		fmt.Println(Token[1])
		h.ServeHTTP(w, r)
	})
}
