package app

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func MySQL() (conn *sql.DB) {
	Driver := "mysql"
	Username := "root"
	Password := "sebastian123_*"
	Name := "mygo"

	conn, err := sql.Open(Driver, Username+":"+Password+Name)

	if err != nil {
		panic(err.Error())
	}
	return conn
}
