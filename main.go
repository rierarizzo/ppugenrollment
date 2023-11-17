package main

import (
	"fmt"
	"ppugenrollment/internal/api"
	"ppugenrollment/internal/data"
	"ppugenrollment/internal/data/user"
	"ppugenrollment/internal/usecases/authenticator"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", "root", "root", "localhost", "3306", "ppugenrollment")
	dbConn := data.ConnectToMySQL(dsn)

	userRepo := user.New(dbConn)
	studentAuthenticator := *authenticator.NewStudentAuthenticator(userRepo)
	adminAuthenticator := *authenticator.NewAdminAuthenticator(userRepo)
	approverAuthenticator := *authenticator.NewApproverAuthenticator(userRepo)

	router := api.Router(studentAuthenticator, adminAuthenticator, approverAuthenticator)

	router.Logger.Fatal(router.Start(fmt.Sprintf(":%s", "8080")))
}
