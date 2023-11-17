package api

import (
	"github.com/labstack/echo/v4"
	"ppugenrollment/internal/api/controllers/authentication"
	"ppugenrollment/internal/usecases/authenticator"
)

func Router(studentAuthenticator authenticator.StudentAuthenticator,
	adminAuthenticator authenticator.AdminAuthenticator,
	approverAuthenticator authenticator.ApproverAuthenticator) *echo.Echo {
	e := echo.New()

	auth := e.Group("/auth")
	authentication.Routes(auth)(studentAuthenticator, adminAuthenticator, approverAuthenticator)

	return e
}
