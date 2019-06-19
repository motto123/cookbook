package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func initMysql() {
	var err error
	//mysqlDsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=%s&multiStatements=true",
	//	setting.LmdbConf.User, setting.LmdbConf.Password, setting.LmdbConf.Host, 3306, setting.LmdbConf.Database, url.QueryEscape(time.Local.String()))
	conn := "root:.1994.10.@tcp(localhost:3306)/cookbook"
	fmt.Println(conn)
	DbApp, err = sql.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	DbApp.SetMaxOpenConns(50000)
	DbApp.SetMaxIdleConns(50000)

	fmt.Println("DbApp database connections pool has built.")
}
