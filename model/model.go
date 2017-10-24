package model

import(
	"fmt"
	"app/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Init() {
	db, err = gorm.Open("mysql", "root:tech0827@/echo_dev?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	} else {
		log.AppLog.Info(fmt.Sprintf("%s", db))
	}
}

func GetConnection() *gorm.DB {
	return db
}
