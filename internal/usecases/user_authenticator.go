package usecases

import (
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"ppugenrollment/internal/domain"
	"ppugenrollment/internal/ports"
	"ppugenrollment/internal/security"
)

type DefaultUserAuthenticator struct {
	userRepo ports.UserRepository
}

func NewUserAuthenticator(userRepo ports.UserRepository) *DefaultUserAuthenticator {
	return &DefaultUserAuthenticator{userRepo}
}

func (a *DefaultUserAuthenticator) Register(user *domain.User) *domain.AppError {
	if err := encryptUserPassword(user); err != nil {
		slog.Error(err.Error())
		return domain.NewAppErrorWithType(domain.UnexpectedError)
	}

	if appErr := a.userRepo.InsertUser(user); appErr != nil {
		slog.Error(appErr.Error())
		return domain.NewAppErrorWithType(domain.UnexpectedError)
	}

	return nil
}

func (a *DefaultUserAuthenticator) Login(email, password string) (*domain.AuthUserPayload, *domain.AppError) {
	user, appErr := a.userRepo.SelectUserByEmail(email)

	if appErr != nil {
		slog.Error(appErr.Error())
		return nil, domain.NewAppErrorWithType(domain.NotAuthenticatedError)
	}

	if err := decryptUserPassword(user.Password, password); err != nil {
		slog.Error(appErr.Error())
		return nil, domain.NewAppErrorWithType(domain.NotAuthenticatedError)
	}

	token, appErr := security.CreateJWTToken(*user)

	if appErr != nil {
		slog.Error(appErr.Error())
		return nil, domain.NewAppErrorWithType(domain.NotAuthenticatedError)
	}

	payload := &domain.AuthUserPayload{
		User:        *user,
		AccessToken: token,
	}

	return payload, nil
}

func encryptUserPassword(user *domain.User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(bytes)
	return nil
}

func decryptUserPassword(userPassword, passwordFromDB string) error {
	return bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(passwordFromDB))
}
