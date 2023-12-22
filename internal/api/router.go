package api

import (
	"github.com/labstack/echo/v4"
	"ppugenrollment/internal/api/controllers"
	"ppugenrollment/internal/api/middlewares"
	"ppugenrollment/internal/usecases"
)

func Router(
	userAuth usecases.DefaultUserAuthenticator,
	projectMngr usecases.DefaultProjectManager,
	enroller usecases.DefaultProjectEnroller,
	approver usecases.DefaultEnrollmentApprover) *echo.Echo {
	e := echo.New()

	middlewares.RoutesAllowedByRoles = routesAllowedByRoles()
	e.Use(middlewares.CORS())

	authRouter(e, userAuth)
	projectManagerRouter(e, projectMngr)
	projectEnrollerRouter(e, enroller)
	approvalRouter(e, approver)

	return e
}

func authRouter(e *echo.Echo, userAuth usecases.DefaultUserAuthenticator) {
	authController := controllers.NewAuthController(&userAuth)

	authGroup := e.Group("/authentication")
	authGroup.POST("/register", authController.Register)
	authGroup.POST("/login", authController.Login)
}

func projectManagerRouter(e *echo.Echo, projectManager usecases.DefaultProjectManager) {
	projectMngrController := controllers.NewProjectController(&projectManager)

	projectMngrGroup := e.Group("/project", middlewares.VerifyJWTAndRoles)
	projectMngrGroup.GET("/getAllProjects", projectMngrController.GetAllProjects)
}

func projectEnrollerRouter(e *echo.Echo, projectEnroller usecases.DefaultProjectEnroller) {
	enrollmentController := controllers.NewEnrollmentController(&projectEnroller)

	enrollmentGroup := e.Group("/enrollment", middlewares.VerifyJWTAndRoles)
	enrollmentGroup.POST("/enrollToProject", enrollmentController.EnrollToProject)
}

func approvalRouter(e *echo.Echo, approver usecases.DefaultEnrollmentApprover) {
	approvalController := controllers.NewApprovalController(&approver)

	approvalGroup := e.Group("/approval", middlewares.VerifyJWTAndRoles)
	approvalGroup.POST("/approveEnrollmentApplication/:application_id", approvalController.ApproveEnrollmentApplication)
}
