package domain

import "time"

type User interface {
	GetEmail() string
	GetPassword() string
}

type CommonUserFields struct {
	ID           int
	IDCardNumber string
	Name         string
	Surname      string
	Email        string
	Password     string
	Role         string
}

type Student struct {
	UserFields  CommonUserFields
	DateOfBirth time.Time
	IsAGraduate bool
	Level       int
}

func (s *Student) GetEmail() string {
	return s.UserFields.Email
}

func (s *Student) GetPassword() string {
	return s.UserFields.Password
}

type Approver struct {
	UserFields CommonUserFields
}

func (a *Approver) GetEmail() string {
	return a.UserFields.Email
}

func (a *Approver) GetPassword() string {
	return a.UserFields.Password
}

type Admin struct {
	UserFields CommonUserFields
}

func (a *Admin) GetEmail() string {
	return a.UserFields.Email
}

func (a *Admin) GetPassword() string {
	return a.UserFields.Password
}

type AuthUserPayload struct {
	UserFields  CommonUserFields
	AccessToken string
}
