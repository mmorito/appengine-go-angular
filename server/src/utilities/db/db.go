package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB
var err error

func Init() *gorm.DB {
	rdbms := "mysql"
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	protocol := os.Getenv("PROTOCOL")
	dbname := os.Getenv("DBNAME")
	option := "charset=utf8&parseTime=True&loc=UTC"

	connect := user + ":" + password + "@" + protocol + "/" + dbname + "?" + option
	db, err = gorm.Open(rdbms, connect)
	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	}

	// SQLログ出力設定
	db.LogMode(true)
	return db
}

func GetConnection() *gorm.DB {
	return db
}
