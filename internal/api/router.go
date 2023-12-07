package api

import (
	"github.com/labstack/echo/v4"
	"ppugenrollment/internal/api/controllers/approval"
	"ppugenrollment/internal/api/controllers/authentication"
	"ppugenrollment/internal/api/controllers/enrollment"
	"ppugenrollment/internal/api/controllers/project"
	"ppugenrollment/internal/usecases/authenticator"
	"ppugenrollment/internal/usecases/enrollment_application_approver"
	"ppugenrollment/internal/usecases/project_enroller"
	"ppugenrollment/internal/usecases/project_manager"
)

func Router(
	userAuth authenticator.DefaultAuthenticator,
	projectMngr project_manager.DefaultManager,
	enroller project_enroller.DefaultEnroller,
	approver enrollment_application_approver.DefaultApprover) *echo.Echo {
	e := echo.New()

	authGroup := e.Group("/auth")
	authentication.Routes(authGroup)(&userAuth)

	projectMngrGroup := e.Group("/projects")
	project.Routes(projectMngrGroup)(&projectMngr)

	enrollmentGroup := e.Group("/enrollment")
	enrollment.Routes(enrollmentGroup)(&enroller)

	approvalGroup := e.Group("/approval")
	approval.Routes(approvalGroup)(&approver)

	return e
}
