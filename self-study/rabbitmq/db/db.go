package db

import (
	"buoi11/util"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/sirupsen/logrus"
)

func Connect() *sql.DB {
	evn, err := util.LoadDBConfig(".")
	if err != nil {
		logrus.Error("Load config fail: ", err)
	}
	conn, err := sql.Open(evn.DBDriver, evn.DSN)
	if err != nil {
		logrus.Error("Open DB fail: ", err)
	}
	return conn
}
