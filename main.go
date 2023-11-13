package main

import (
	"fmt"
	"ppugenrollment/internal/api"
	"ppugenrollment/internal/data"
	"ppugenrollment/internal/data/student"
	"ppugenrollment/internal/usecases/student_authenticator"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", "root", "root", "localhost", "3306", "ppugenrollment")
	dbConn := data.ConnectToMySQL(dsn)

	studentRepo := student.New(dbConn)
	defaultStudentAuthenticator := student_authenticator.New(studentRepo)

	router := api.Router(defaultStudentAuthenticator)

	router.Logger.Fatal(router.Start(fmt.Sprintf(":%s", "8080")))
}
