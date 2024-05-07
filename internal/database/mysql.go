package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Database_Conn *sql.DB
var err error

func ConnectMysqlDB() error {
	Database_Conn, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/employee_management")
	if err != nil {
		return err
	}
	if err = Database_Conn.Ping(); err != nil {
		return err
	}
	Database_Conn.SetMaxOpenConns(5)
	Database_Conn.SetMaxIdleConns(5)
	return nil
}

func DisConnectMysqlDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		fmt.Println(err)
	}
}
