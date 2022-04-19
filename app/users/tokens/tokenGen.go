package tokens

import (
	"math/rand"
)

var basicString string = `aZbYcX%9WdVe@UfT8S#gRhQi[P7OjN]kMlL6$KmJnI!oH5Gp?FqErD4CsBtAu3vwq2yz1`

// Генерирует врменный токен
func GenerateTempToken() string {
	token := ""
	for i := 0; i <= 10; i++ {
		token += string(basicString[rand.Intn(len(basicString))])
	}
	return token
}

// Генерирует постоянный токен
func GenerateMainToken() string {
	token := ""
	for i := 0; i <= 15; i++ {
		token += string(basicString[rand.Intn(len(basicString))])
	}
	return token
}
