package ports

import "ppugenrollment/pkg/domain"

type UserAuthenticator interface {
	Register(userRegistrable *domain.User) *domain.AppError
	Login(email, password string) (*domain.AuthUserPayload, *domain.AppError)
}

type UserRepository interface {
	InsertUser(student *domain.User) *domain.AppError
	SelectUserByEmail(email string) (*domain.User, *domain.AppError)
}
