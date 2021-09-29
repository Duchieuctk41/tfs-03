package database

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

type MysqlDb struct {
	SQL *sql.DB
}

var Mysql = &MysqlDb{}

func Connect(user, password, host, port, dbname string) *MysqlDb {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		user, password, host, port, dbname)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
	Mysql.SQL = db
	return Mysql
}
