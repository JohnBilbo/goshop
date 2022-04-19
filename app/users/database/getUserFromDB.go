package database

import (
	"bshop/app/users/models"
	"context"

	"github.com/jackc/pgx/v4"
)

// Выдает пользователя из базы данных.
// Если пользователя нет то выдет ошибку
func GetUserFromDB(username, password, email string) (*models.RegistrationData, error) {
	db, err := pgx.Connect(context.Background(), GetDBConfig())
	if err != nil {
		return nil, err
	}
	defer db.Close(context.Background())
	sqlString := `SELECT id, username, password, token_temp FROM shop_user WHERE username=$1;`

	queryset, err := db.Query(context.Background(), sqlString, username)
	if err != nil {
		return nil, err
	}
	userData := models.RegistrationData{}
	for queryset.Next() {
		queryset.Scan(
			&userData.Id,
			&userData.Username,
			&userData.Password,
			&userData.TokenTemp,
		)
	}
	return &userData, nil
}
