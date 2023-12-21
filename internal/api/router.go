package api

import (
	"github.com/labstack/echo/v4"
	"ppugenrollment/internal/api/controllers"
	"ppugenrollment/internal/api/middlewares"
	"ppugenrollment/internal/usecases"
)

func Router(
	userAuth usecases.DefaultAuthenticator,
	projectMngr usecases.DefaultManager,
	enroller usecases.DefaultEnroller,
	approver usecases.DefaultApprover) *echo.Echo {
	e := echo.New()

	e.Use(middlewares.Middleware())

	authGroup := e.Group("/authentication")
	controllers.AuthRoutes(authGroup)(&userAuth)

	middlewares.RoutesAllowedByRoles = routesAllowedByRoles()

	projectMngrGroup := e.Group("/project", middlewares.Verify)
	controllers.ProjectRoutes(projectMngrGroup)(&projectMngr)

	enrollmentGroup := e.Group("/enrollment", middlewares.Verify)
	controllers.EnrollmentRoutes(enrollmentGroup)(&enroller)

	approvalGroup := e.Group("/approval", middlewares.Verify)
	controllers.ApprovalRoutes(approvalGroup)(&approver)

	return e
}

func routesAllowedByRoles() map[string][]string {
	return map[string][]string{
		"/authentication": {"ALL"},
		"/project":        {"ALL"},
		"/enrollment":     {"S"},
		"/approval":       {"A"},
	}
}
