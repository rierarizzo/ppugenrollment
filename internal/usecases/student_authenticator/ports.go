package student_authenticator

import (
	"ppugenrollment/internal/domain"
)

type Authenticator interface {
	Register(student *domain.Student) *domain.AppError
	Login(email, password string) (*domain.AuthenticatedUser, *domain.AppError)
}

type UserRepository interface {
	InsertStudent(student *domain.Student) *domain.AppError
	GetStudentByEmail(email string) (*domain.Student, *domain.AppError)
}
