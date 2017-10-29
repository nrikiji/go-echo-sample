package db

import(
	"app/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Init() {
	
	c := config.Config.Database
	
	user := c.User
	password := c.Password
	dbname := c.Name
	
	db, err = gorm.Open("mysql", user + ":" + password + "@/" + dbname + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	}
}

func GetConnection() *gorm.DB {
	return db
}
