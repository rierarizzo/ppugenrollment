package api

import (
	"github.com/labstack/echo/v4"
	"ppugenrollment/internal/api/controllers"
	"ppugenrollment/internal/api/middlewares/auth"
	"ppugenrollment/internal/api/middlewares/cors"
	"ppugenrollment/internal/usecases"
)

func Router(
	userAuth usecases.DefaultAuthenticator,
	projectMngr usecases.DefaultManager,
	enroller usecases.DefaultEnroller,
	approver usecases.DefaultApprover) *echo.Echo {
	e := echo.New()

	e.Use(cors.Middleware())

	authGroup := e.Group("/authentication")
	controllers.AuthRoutes(authGroup)(&userAuth)

	auth.RoutesAllowedByRoles = routesAllowedByRoles()

	projectMngrGroup := e.Group("/project", auth.Verify)
	controllers.ProjectRoutes(projectMngrGroup)(&projectMngr)

	enrollmentGroup := e.Group("/enrollment", auth.Verify)
	controllers.EnrollmentRoutes(enrollmentGroup)(&enroller)

	approvalGroup := e.Group("/approval", auth.Verify)
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
