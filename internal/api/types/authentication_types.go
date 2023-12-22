package types

import (
	"ppugenrollment/internal/api/utils"
	"ppugenrollment/internal/domain"
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
	v := new(utils.Validator)

	v.MustNotBeEmpty(r.IDCardNumber)
	v.MustNotBeEmpty(r.Name)
	v.MustNotBeEmpty(r.Surname)
	v.MustNotBeEmpty(r.Email)
	v.MustBeEmail(r.Email)
	v.MustNotBeEmpty(r.Password)
	v.MustNotBeEmpty(r.Role)

	return v.AppErr
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *UserLoginRequest) Validate() *domain.AppError {
	v := new(utils.Validator)

	v.MustNotBeEmpty(r.Email)
	v.MustBeEmail(r.Email)
	v.MustNotBeEmpty(r.Password)

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
