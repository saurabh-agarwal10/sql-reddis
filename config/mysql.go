package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var MysqlConnection *sql.DB
var err error

func ConnectMysql() error {
	if MysqlConnection != nil {
		return nil
	}

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASS"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DBNAME"))

	
	MysqlConnection, err = sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	log.Printf("MySQL Successfully Connected.")
	return nil
}
