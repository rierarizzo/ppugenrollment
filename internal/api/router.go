package api

import (
	"ppugenrollment/internal/api/controllers"
	"ppugenrollment/internal/api/middlewares"
	"ppugenrollment/internal/ports"

	"github.com/labstack/echo/v4"
)

func Router(userAuth ports.UserAuthenticator, projectMngr ports.ProjectManager, enroller ports.ProjectEnroller,
	approver ports.EnrollmentApprover) *echo.Echo {
	e := echo.New()

	loadMiddlewares(e)

	loadAuthRouter(e, userAuth)
	loadProjectManagerRouter(e, projectMngr)
	loadProjectEnrollerRouter(e, enroller)
	loadApprovalRouter(e, approver)

	return e
}

func loadAuthRouter(e *echo.Echo, userAuth ports.UserAuthenticator) {
	authController := controllers.NewAuthController(userAuth)

	authGroup := e.Group("/authentication")
	authGroup.POST("/register", authController.Register)
	authGroup.POST("/login", authController.Login)
}

func loadProjectManagerRouter(e *echo.Echo, projectManager ports.ProjectManager) {
	projectMngrController := controllers.NewProjectController(projectManager)

	projectMngrGroup := e.Group("/project", middlewares.VerifyJWTAndRoles)
	projectMngrGroup.GET("/getAllProjects", projectMngrController.GetAllProjects)
	projectMngrGroup.GET("/getProjectById/:project_id", projectMngrController.GetProjectByID)
	projectMngrGroup.POST("/insertNewProject", projectMngrController.AddNewProject)
	projectMngrGroup.PUT("/updateProject/:project_id", projectMngrController.UpdateProject)
	projectMngrGroup.DELETE("/deleteProject/:project_id", projectMngrController.DeleteProject)
	projectMngrGroup.GET("/getCompanies", projectMngrController.GetCompanies)
}

func loadProjectEnrollerRouter(e *echo.Echo, projectEnroller ports.ProjectEnroller) {
	enrollmentController := controllers.NewEnrollmentController(projectEnroller)

	enrollmentGroup := e.Group("/enrollment", middlewares.VerifyJWTAndRoles)
	enrollmentGroup.POST("/enrollToProject", enrollmentController.EnrollToProject)
	enrollmentGroup.GET("/getEnrollmentApplications", enrollmentController.GetEnrollmentApplications)
}

func loadApprovalRouter(e *echo.Echo, approver ports.EnrollmentApprover) {
	approvalController := controllers.NewApprovalController(approver)

	approvalGroup := e.Group("/approval", middlewares.VerifyJWTAndRoles)
	approvalGroup.POST("/approveEnrollmentApplication/:application_id", approvalController.ApproveEnrollmentApplication)
}

func loadMiddlewares(e *echo.Echo) {
	middlewares.RoutesAllowedByRoles = routesAllowedByRoles()
	e.Use(middlewares.CORS())
}
