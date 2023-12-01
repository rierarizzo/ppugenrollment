package api

import (
	"github.com/labstack/echo/v4"
	"ppugenrollment/internal/api/controllers/authentication"
	"ppugenrollment/internal/usecases/authenticator"
)

func Router(userAuth authenticator.UserAuthenticator) *echo.Echo {
	e := echo.New()

	auth := e.Group("/auth")
	authentication.Routes(auth)(userAuth)

	return e
}
