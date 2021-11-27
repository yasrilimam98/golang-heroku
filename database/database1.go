package database

import (
	"database/sql"

	"github.com/yasrilimam98/grb-restapi/config"
)

var db1 *sql.DB
var err1 error

func Initdb1() {
	confdb2 := config.GetConfigdb1()
	connetionStringdb2 := confdb2.DB_USERNAME + ":" + confdb2.DB_PASSWORD + "@tcp(" + confdb2.DB_HOST + ":" + confdb2.DB_PORT + ")/" + confdb2.DB_NAME

	db1, err1 = sql.Open("mysql", connetionStringdb2)

	if err1 != nil {
		panic("Connection String error!")
	}

	err1 = db1.Ping()
	if err1 != nil {
		panic("DSN Invalid")
	}
}

func CreateCondb1() *sql.DB {
	return db1
}
