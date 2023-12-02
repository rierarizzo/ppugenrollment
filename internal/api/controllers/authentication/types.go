package authentication

import "time"

// Requests

type UserRequest struct {
	IDCardNumber string    `json:"id_card_number,omitempty"`
	Name         string    `json:"name,omitempty"`
	Surname      string    `json:"surname,omitempty"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Role         string    `json:"role,omitempty"`
	DateOfBirth  time.Time `json:"date_of_birth,omitempty"`
	IsAGraduate  bool      `json:"is_a_graduate,omitempty"`
	Level        int       `json:"level,omitempty"`
}

// Responses

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
