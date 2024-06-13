package db

import (
	"time"

	"github.com/donaldgifford/tbccgolf/loggy"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	database *gorm.DB
	e        error
)

const (
	DBFile = "tbcc.db"
)

func Init() {
	newLoggers := logger.New(
		loggy.Loggy(),
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL Threshold
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)

	database, e = gorm.Open(sqlite.Open(DBFile), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
		Logger:                   newLoggers,
		SkipDefaultTransaction:   true,
		DisableNestedTransaction: true,
	})
	if e != nil {
		panic(e)
	}
}

func DB() *gorm.DB {
	return database
}
