package authenticator

import (
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"ppugenrollment/internal/domain"
	"ppugenrollment/internal/security"
)

func loginInSystem(email, password string, userRepo UserRepository) (*domain.AuthUserPayload, *domain.AppError) {
	notAuthErr := domain.NewAppErrorWithType(domain.NotAuthenticatedError)

	user, appErr := userRepo.SelectUserByEmail(email)
	if appErr != nil {
		slog.Error(appErr.Error())
		return nil, notAuthErr
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		slog.Error(err.Error())
		return nil, notAuthErr
	}

	token, appErr := security.CreateJWTToken(*user)
	if appErr != nil {
		slog.Error(appErr.Error())
		return nil, notAuthErr
	}

	return &domain.AuthUserPayload{
		UserFields:  *user,
		AccessToken: token,
	}, nil
}
