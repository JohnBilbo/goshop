package database

import (
	"bshop/app/users/models"
	"bshop/app/users/tokens"
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// Регистрирует пользователя в БД
func UserRegistrationDB(userData *models.RegistrationData) (pgconn.CommandTag, error) {
	userData.TokenMain = tokens.GenerateMainToken()
	userData.TokenTemp = tokens.GenerateTempToken()
	userData.Balance = 0
	sqlString := `INSERT INTO 
		shop_user(username, password, email, token_main, token_temp, balance) 
		VALUES($1,$2,$3,$4,$5,$6);`
	db, err := pgx.Connect(context.Background(), GetDBConfig())
	if err != nil {
		return nil, err
	}
	defer db.Close(context.Background())
	commandTag, err := db.Exec(
		context.Background(),
		sqlString,
		userData.Username,
		userData.Password,
		userData.Email,
		userData.TokenMain,
		userData.TokenTemp,
		userData.Balance)
	if err != nil {
		return nil, err
	}
	return commandTag, nil
}
