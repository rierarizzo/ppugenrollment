package types

import (
	"ppugenrollment/pkg/domain"
	"time"
)

type UserRegisterRequest struct {
	IDCardNumber string    `json:"id_card_number"`
	Name         string    `json:"name"`
	Surname      string    `json:"surname"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Role         string    `json:"role"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	IsAGraduate  bool      `json:"is_a_graduate"`
	Level        int       `json:"level"`
}

func (r *UserRegisterRequest) Validate() *domain.AppError {
	v := new(Validator)

	v.MustNotBeEmptyString(r.IDCardNumber)
	v.MustNotBeEmptyString(r.Name)
	v.MustNotBeEmptyString(r.Surname)
	v.MustNotBeEmptyString(r.Email)
	v.MustBeEmail(r.Email)
	v.MustNotBeEmptyString(r.Password)
	v.MustNotBeEmptyString(r.Role)

	return v.AppErr
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *UserLoginRequest) Validate() *domain.AppError {
	v := new(Validator)

	v.MustNotBeEmptyString(r.Email)
	v.MustBeEmail(r.Email)
	v.MustNotBeEmptyString(r.Password)

	return v.AppErr
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
