package dao_test

import (
	"log"
	"os"
	"tiktok/configs"
	"tiktok/dao"
	"tiktok/middleware/redis"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var testDb *gorm.DB

func InitTestDb() error {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	var err error
	testDb, err = gorm.Open(mysql.Open(configs.GetTestDBInfo()), &gorm.Config{
		Logger:                 newLogger,
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Panicln("err :", err.Error())
		return err
	}
	return nil
}

func InitTestRedis() {
	redis.Init()
}

func TestInit() error {
	if err := InitTestDb(); err != nil {
		return err
	}
	InitTestRedis()
	dao.DB = testDb
	return nil
}
