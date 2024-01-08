package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectToMySQL(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db
}
