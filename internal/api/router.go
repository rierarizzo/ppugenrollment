package api

import (
	"github.com/labstack/echo/v4"
	"ppugenrollment/internal/api/controllers/authentication"
	"ppugenrollment/internal/usecases/student_authenticator"
)

func Router(studentAuthenticator student_authenticator.Authenticator) *echo.Echo {
	e := echo.New()

	auth := e.Group("/auth")
	authentication.Routes(auth)(studentAuthenticator)

	return e
}
