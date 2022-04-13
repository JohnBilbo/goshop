package controllers

import (
	"bshop/app/users/database"
	"bshop/app/users/models"
	"encoding/json"
	"fmt"
	"net/http"
)

// Регистрирует пользователя
func UserRegistrationController(w http.ResponseWriter, r *http.Request) {
	var userData models.RegistrationData
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		fmt.Println("JSON Decode error ", err)
		http.Error(w, "Incorrect data", 500)
		return
	}
	if userData.Username == "" || userData.Password == "" || userData.Email == "" {
		fmt.Println("Empty DATA Error ", err)
		http.Error(w, "Empty data", 500)
		return
	}
	commandTag, err := database.UserRegistrationDB(&userData)
	// Заметка Нужно правильно обработать ошибки БД отдаваемые в Response
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	if commandTag != nil {
		w.Write([]byte("User added\n"))
	}
}

// Обновить данные пользователя
func UserUpdateController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Worked\n"))
}

// Удалить пользователя
func UserRemoveController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Worked\n"))
}
