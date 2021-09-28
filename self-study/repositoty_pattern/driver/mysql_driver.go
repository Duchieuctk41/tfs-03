package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlDB struct {
	SQL *sql.DB
}

var Mysql = &MySqlDB{}

func Connect(host, port, user, password, dbname string) *MySqlDB {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		user, password, host, port, dbname)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
	Mysql.SQL = db
	return Mysql
}
