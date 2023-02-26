package database

import (
	"database/sql"
	"github.com/ffmoyano/gofer/logger"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

var db *sql.DB

func Open() {
	var err error
	driver := os.Getenv("dbDriver")
	connString := os.Getenv("dbUrl")
	db, err = sql.Open(driver, connString)
	if err != nil {
		logger.Error(err.Error())
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxIdleTime(60 * time.Second)
	db.SetConnMaxLifetime(30 * time.Minute)
}
