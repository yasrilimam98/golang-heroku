package database

import (
	"database/sql"

	"github.com/yasrilimam98/grb-restapi/config"
)

var db2 *sql.DB
var err2 error

func Initdb2() {
	confdb2 := config.GetConfigdb2()
	connetionStringdb2 := confdb2.DB_USERNAME + ":" + confdb2.DB_PASSWORD + "@tcp(" + confdb2.DB_HOST + ":" + confdb2.DB_PORT + ")/" + confdb2.DB_NAME

	db2, err2 = sql.Open("mysql", connetionStringdb2)

	if err2 != nil {
		panic("Connection String error!")
	}

	err2 = db2.Ping()
	if err2 != nil {
		panic("DSN Invalid")
	}
}

func CreateCondb2() *sql.DB {
	return db2
}
