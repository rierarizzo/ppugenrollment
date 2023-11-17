package security

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"os"
	"ppugenrollment/internal/domain"
	"time"
)

const SecretKey = "SECRET_KEY"

func CreateJWTToken(user domain.CommonUserFields) (string, *domain.AppError) {
	secret := []byte(os.Getenv(SecretKey))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{
		Id:    user.ID,
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Hour),
			},
		},
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", domain.NewAppError(err, domain.TokenGenerationError)
	}

	return tokenString, nil
}

func VerifyJWTToken(tokenString string) (*UserClaims, error) {
	secret := []byte(os.Getenv(SecretKey))

	var userClaims UserClaims
	token, err := jwt.ParseWithClaims(tokenString, &userClaims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.NewAppErrorWithType(domain.TokenValidationError)
		}

		return secret, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			slog.Error("Token is expired")
		}

		return nil, domain.NewAppError(err, domain.TokenValidationError)
	}
	if !token.Valid {
		return nil, domain.NewAppError(err, domain.TokenValidationError)
	}

	return &userClaims, nil
}
