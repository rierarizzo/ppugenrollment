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
	userAuth := *authenticator.NewUserAuthenticator(userRepo)

	router := api.Router(userAuth)

	router.Logger.Fatal(router.Start(fmt.Sprintf(":%s", "8080")))
}
