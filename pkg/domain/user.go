package domain

import "time"

type User struct {
	ID           int
	IDCardNumber string
	Name         string
	Surname      string
	Email        string
	Password     string
	Role         string
	DateOfBirth  time.Time
	IsAGraduate  bool
	Level        int
}

type AuthUserPayload struct {
	User
	AccessToken string
}
