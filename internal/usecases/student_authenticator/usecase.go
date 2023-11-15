package student_authenticator

import (
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"ppugenrollment/internal/domain"
	"ppugenrollment/internal/security"
)

type DefaultAuthenticator struct {
	userRepository UserRepository
}

func New(userRepository UserRepository) *DefaultAuthenticator {
	return &DefaultAuthenticator{userRepository: userRepository}
}

func (d DefaultAuthenticator) Register(student *domain.Student) *domain.AppError {
	bytes, err := bcrypt.GenerateFromPassword([]byte(student.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.NewAppError(err, domain.UnexpectedError)
	}

	student.User.Password = string(bytes)

	return d.userRepository.InsertStudent(student)
}

func (d DefaultAuthenticator) Login(email, password string) (*domain.AuthenticatedUser, *domain.AppError) {
	student, appErr := d.userRepository.GetStudentByEmail(email)
	if appErr != nil {
		slog.Error(appErr.Error())
		return nil, domain.NewAppErrorWithType(domain.NotAuthenticatedError)
	}

	err := bcrypt.CompareHashAndPassword([]byte(student.User.Password), []byte(password))
	if err != nil {
		slog.Error(appErr.Error())
		return nil, domain.NewAppErrorWithType(domain.NotAuthenticatedError)
	}

	token, appErr := security.CreateJWTToken(student.User)
	if appErr != nil {
		return nil, appErr
	}

	authenticatedUser := &domain.AuthenticatedUser{
		User:        student.User,
		AccessToken: token,
	}

	return authenticatedUser, nil
}
