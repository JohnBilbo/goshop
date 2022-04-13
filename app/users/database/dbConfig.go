package database

import (
	"fmt"
)

// Возвращает конфигурационные дынные для подключения к БД в виде строки
func GetDBConfig() string {
	host := "127.0.0.1"
	port := "5432"
	user := "goshop_user"
	password := "4554"
	dbname := "goshop"
	connectionConfig := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	return connectionConfig
}
