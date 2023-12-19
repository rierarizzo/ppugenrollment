package main

import (
	"fmt"
	"os"
	"ppugenrollment/internal/api"
	"ppugenrollment/internal/data"
	"ppugenrollment/internal/data/approval"
	"ppugenrollment/internal/data/enrollment"
	"ppugenrollment/internal/data/project"
	"ppugenrollment/internal/data/user"
	"ppugenrollment/internal/usecases/enrollment_application_approver"
	"ppugenrollment/internal/usecases/project_enroller"
	"ppugenrollment/internal/usecases/project_manager"
	"ppugenrollment/internal/usecases/user_authenticator"
)

const webPort = "80"

func main() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))
	dbConn := data.ConnectToMySQL(dsn)

	userRepo := user.New(dbConn)
	userAuth := *user_authenticator.New(userRepo)

	projectRepo := project.New(dbConn)
	projectMngr := *project_manager.New(projectRepo)

	enrollmentRepo := enrollment.New(dbConn)
	projectEnroller := *project_enroller.New(enrollmentRepo)

	approvalRepo := approval.New(dbConn)
	applicationApprover := *enrollment_application_approver.New(approvalRepo)

	router := api.Router(userAuth, projectMngr, projectEnroller, applicationApprover)

	router.Logger.Fatal(router.Start(fmt.Sprintf(":%s", webPort)))
}
