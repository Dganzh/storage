package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/fileserver?charset=utf8")
	db.SetMaxOpenConns(100)
	if err := db.Ping(); err != nil {
		fmt.Println("failed connect to mysql, err: " + err.Error())
		os.Exit(1)
	}
}

func DBConn() *sql.DB {
	return db
}



