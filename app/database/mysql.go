package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func MySQL() (conn *sql.DB) {
	Driver := "mysql"
	Username := "root"
	Password := "sebastian123_*"
	Name := "mygo"

	connStr := fmt.Sprintf("%s:%s@/%s", Username, Password, Name)
	conn, err := sql.Open(Driver, connStr)

	if err != nil {
		panic(err.Error())
	}
	return conn
}
