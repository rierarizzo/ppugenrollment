package types

import "time"

type UserRegisterRequest struct {
	IDCardNumber string    `json:"id_card_number" validate:"required"`
	Name         string    `json:"name" validate:"required"`
	Surname      string    `json:"surname" validate:"required"`
	Email        string    `json:"email" validate:"required,email"`
	Password     string    `json:"password" validate:"required"`
	Role         string    `json:"role" validate:"required"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	IsAGraduate  bool      `json:"is_a_graduate"`
	Level        int       `json:"level"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	Surname              string `json:"surname"`
	Email                string `json:"email"`
	Role                 string `json:"role"`
	AuthenticationResult struct {
		AccessToken string `json:"access_token"`
	} `json:"authentication_result,omitempty"`
}
