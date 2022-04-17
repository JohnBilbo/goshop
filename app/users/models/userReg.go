package models

// shop_user table
type RegistrationData struct {
	Id        int     `json:"id" db:"id"`
	Username  string  `json:"username" db:"username"`
	Password  string  `json:"password" db:"password"`
	Email     string  `json:"email" db:"email"`
	TokenMain string  `json:"token_main" db:"token_main"`
	TokenTemp string  `json:"token_temp" db:"token_temp"`
	Balance   float64 `json:"balance" db:"balance"`
}
