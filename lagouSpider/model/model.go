package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB

	username string = "remote"
	password string = "123456"
	dbName   string = "spiders"
)

func init() {
	var err error
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(192.168.2.30:3306)/%s?charset=utf8&parseTime=True", username, password, dbName))
	if err != nil {
		log.Fatalf(" gorm.Open.err: %v", err)
	}

	DB.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}
}
