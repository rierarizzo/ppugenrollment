package user_authenticator

import (
	"ppugenrollment/internal/domain"
)

type Authenticator interface {
	Register(userRegistrable domain.UserRegistrable) *domain.AppError
	Login(email, password string) (*domain.AuthUserPayload, *domain.AppError)
}

type UserRepository interface {
	InsertStudent(student *domain.Student) *domain.AppError
	InsertAdmin(admin *domain.Admin) *domain.AppError
	InsertApprover(approver *domain.Approver) *domain.AppError

	SelectUserByEmail(email string) (*domain.User, *domain.AppError)
}
