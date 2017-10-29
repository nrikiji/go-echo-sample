package db

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Campaign struct {
	gorm.Model
	Name string
}
