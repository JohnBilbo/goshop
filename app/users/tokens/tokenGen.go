package tokens

import (
	"math/rand"
)

// Генерирует врменный токен
func GenerateTempToken() string {
	basicString := `aZbYcX%9WdVe@UfT8S#gRhQi[P7OjN]kMlL6$KmJnI!oH5Gp?FqErD4CsBtAu3vwq2yz1`
	token := ""
	for i := 0; i <= 15; i++ {
		token += string(basicString[rand.Intn(15)+rand.Intn(15)])
	}
	return token
}

// Генерирует постоянный токен
func GenerateMainToken() string {
	basicString := `aZbYcX%9WdVe@UfT8S#gRhQi[P7OjN]kMlL6$KmJnI!oH5Gp?FqErD4CsBtAu3vwq2yz1`
	token := ""
	for i := 0; i <= 20; i++ {
		token += string(basicString[rand.Intn(20)+rand.Intn(20)])
	}
	return token
}
