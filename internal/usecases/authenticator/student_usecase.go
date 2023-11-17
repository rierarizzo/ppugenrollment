package authenticator

import (
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"ppugenrollment/internal/domain"
)

type StudentAuthenticator struct {
	userRepo UserRepository
}

func NewStudentAuthenticator(userRepo UserRepository) *StudentAuthenticator {
	return &StudentAuthenticator{userRepo}
}

func (d *StudentAuthenticator) Register(user domain.User) *domain.AppError {
	var student *domain.Student
	if s, ok := user.(*domain.Student); ok {
		student = s
	} else {
		return domain.NewAppErrorWithType(domain.UnexpectedError)
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(student.UserFields.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error(err.Error())
		return domain.NewAppError(err, domain.UnexpectedError)
	}

	student.UserFields.Password = string(bytes)

	appErr := d.userRepo.InsertStudent(student)
	if appErr != nil {
		slog.Error(appErr.Error())
		return domain.NewAppError(appErr, domain.UnexpectedError)
	}

	return nil
}

func (d *StudentAuthenticator) Login(email, password string) (*domain.AuthUserPayload, *domain.AppError) {
	return loginInSystem(email, password, d.userRepo)
}
