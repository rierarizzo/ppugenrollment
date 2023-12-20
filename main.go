package main

import (
	"fmt"
	"os"
	"ppugenrollment/internal/api"
	"ppugenrollment/internal/data"
	"ppugenrollment/internal/data/repository"
	"ppugenrollment/internal/usecases"
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

	userRepo := repository.NewUserRepository(dbConn)
	userAuth := *usecases.NewUserAuthenticator(userRepo)

	projectRepo := repository.NewProjectRepository(dbConn)
	projectMngr := *usecases.NewProjectManager(projectRepo)

	enrollmentRepo := repository.NewEnrollmentRepository(dbConn)
	projectEnroller := *usecases.NewProjectEnroller(enrollmentRepo)

	approvalRepo := repository.NewApprovalRepository(dbConn)
	applicationApprover := *usecases.NewEnrollmentApprover(approvalRepo)

	router := api.Router(userAuth, projectMngr, projectEnroller, applicationApprover)

	router.Logger.Fatal(router.Start(fmt.Sprintf(":%s", webPort)))
}
