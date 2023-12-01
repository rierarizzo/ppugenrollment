package domain

import "time"

type UserRegistrable interface {
	GetUser() *User
}

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
	User
	DateOfBirth time.Time
	IsAGraduate bool
	Level       int
}

func (s *Student) GetUser() *User {
	return &s.User
}

type Approver struct {
	User
}

func (a *Approver) GetUser() *User {
	return &a.User
}

type Admin struct {
	User
}

func (a *Admin) GetUser() *User {
	return &a.User
}

type AuthUserPayload struct {
	User
	AccessToken string
}
