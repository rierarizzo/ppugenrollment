package data

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectToMySQL(dsn string) *sqlx.DB {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
