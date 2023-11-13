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
}

type Student struct {
	User        User
	DateOfBirth time.Time
	IsAGraduate bool
	Level       int
}

type Approver struct {
	User User
}

type Admin struct {
	User User
}

type AuthenticatedUser struct {
	User        User
	AccessToken string
}
