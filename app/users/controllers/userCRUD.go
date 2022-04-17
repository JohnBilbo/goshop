package controllers

import (
	"bshop/app/users/database"
	"bshop/app/users/models"
	"encoding/json"
	"net/http"
)

// Регистрирует пользователя
func UserRegistrationController(w http.ResponseWriter, r *http.Request) {
	var userData models.RegistrationData
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", 500)
		return
	}
	if userData.Username == "" || userData.Password == "" || userData.Email == "" {
		http.Error(w, "Не верно заполнены данные", 500)
		return
	}
	commandTag, err := database.UserRegistrationDB(&userData)
	// Заметка - Нужно правильно обработать ошибки БД отдаваемые в Response
	if err != nil && err.Error() == `ОШИБКА: повторяющееся значение ключа нарушает ограничение уникальности "shop_user_username_key" (SQLSTATE 23505)` {
		http.Error(w, "Пользователь с такими данными уже существует\n", 500)
	}
	if commandTag != nil {
		w.Write([]byte("Пользователь добавлен"))
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

// Возвращает доп данные о пользователе
func UserGetController(w http.ResponseWriter, r *http.Request) {
}
