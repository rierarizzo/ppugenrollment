package api

import (
	"github.com/labstack/echo/v4"
	"ppugenrollment/internal/api/controllers/authentication"
	"ppugenrollment/internal/api/controllers/project"
	"ppugenrollment/internal/usecases/authenticator"
	"ppugenrollment/internal/usecases/project_manager"
)

func Router(userAuth authenticator.UserAuthenticator, projectMngr project_manager.DefaultManager) *echo.Echo {
	e := echo.New()

	authGroup := e.Group("/auth")
	authentication.Routes(authGroup)(userAuth)

	projectMngrGroup := e.Group("/projects")
	project.Routes(projectMngrGroup)(projectMngr)

	return e
}
