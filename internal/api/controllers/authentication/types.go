package authentication

import "time"

type StudentRequest struct {
	CommonFields
	DateOfBirth time.Time `json:"date_of_birth"`
	IsAGraduate bool      `json:"is_a_graduate"`
	Level       int       `json:"level"`
}

type AdminRequest struct {
	CommonFields
}

type ApproverRequest struct {
	CommonFields
}

type CommonFields struct {
	IDCardNumber string `json:"id_card_number"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Role         string `json:"role"`
}
