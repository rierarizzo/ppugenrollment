package usecases

import (
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"ppugenrollment/internal/domain"
	"ppugenrollment/internal/ports"
	"ppugenrollment/internal/security"
)

type DefaultAuthenticator struct {
	userRepo ports.UserRepository
}

func NewUserAuthenticator(userRepo ports.UserRepository) *DefaultAuthenticator {
	return &DefaultAuthenticator{userRepo}
}

func (a *DefaultAuthenticator) Register(user *domain.User) *domain.AppError {
	appErr := encryptUserPassword(user)

	if appErr != nil {
		return handleError(appErr, domain.UnexpectedError)
	}

	appErr = a.userRepo.InsertUser(user)

	if appErr != nil {
		return handleError(appErr, domain.UnexpectedError)
	}

	return nil
}

func (a *DefaultAuthenticator) Login(email, password string) (*domain.AuthUserPayload, *domain.AppError) {
	user, appErr := a.userRepo.SelectUserByEmail(email)

	if appErr != nil {
		return nil, handleError(appErr, domain.NotAuthenticatedError)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, handleError(err, domain.NotAuthenticatedError)
	}

	token, appErr := security.CreateJWTToken(*user)

	if appErr != nil {
		return nil, handleError(appErr, domain.NotAuthenticatedError)
	}

	payload := &domain.AuthUserPayload{
		User:        *user,
		AccessToken: token,
	}

	return payload, nil
}

func encryptUserPassword(user *domain.User) *domain.AppError {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return handleError(err, domain.UnexpectedError)
	}

	user.Password = string(bytes)
	return nil
}

func handleError(err error, errType string) *domain.AppError {
	slog.Error(err.Error())
	return domain.NewAppError(err, errType)
}
