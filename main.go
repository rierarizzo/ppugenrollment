package main

import (
	"fmt"
	"ppugenrollment/internal/api"
	"ppugenrollment/internal/data"
	"ppugenrollment/internal/data/project"
	"ppugenrollment/internal/data/user"
	"ppugenrollment/internal/usecases/authenticator"
	"ppugenrollment/internal/usecases/project_manager"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", "root", "root", "localhost", "3306", "ppugenrollment")
	dbConn := data.ConnectToMySQL(dsn)

	userRepo := user.New(dbConn)
	userAuth := *authenticator.New(userRepo)

	projectRepo := project.New(dbConn)
	projectMngr := *project_manager.New(projectRepo)

	router := api.Router(userAuth, projectMngr)

	router.Logger.Fatal(router.Start(fmt.Sprintf(":%s", "8080")))
}
