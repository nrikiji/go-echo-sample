package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	"go-echo-starter/env"

	"github.com/go-testfixtures/testfixtures/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(e env.Env) *gorm.DB {
	return NewDB(e)
}

func NewDB(e env.Env) *gorm.DB {
	db, err := gorm.Open(mysql.Open(e.Dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			loggerConfig(e),
		),
	})
	if err != nil {
		log.Fatalln(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}

	sqlDB.SetMaxIdleConns(3)
	return db
}

func loggerConfig(e env.Env) logger.Config {
	if e.Stage == "production" {
		return loggerProdConfig()
	} else {
		return loggerDevConfig()
	}
}

func loggerDevConfig() logger.Config {
	return logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  logger.Info,
		Colorful:                  true,
		IgnoreRecordNotFoundError: false,
	}
}

func loggerProdConfig() logger.Config {
	return logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  logger.Error,
		Colorful:                  true,
		IgnoreRecordNotFoundError: true,
	}
}

func NewFixtures(s env.Env) *testfixtures.Loader {
	db, err := sql.Open("mysql", s.Dsn)

	if err != nil {
		log.Fatalln(err)
	}

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory("./fixtures"),
	)

	if err != nil {
		log.Fatalln(err)
	}

	return fixtures
}
