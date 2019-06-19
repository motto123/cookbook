package db

import (
	"database/sql"
)

var (
	DbApp   *sql.DB
)

func InitDB() {
	initMysql()
}
