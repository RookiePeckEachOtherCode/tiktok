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

func InitTestDb() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	testDb, _ = gorm.Open(mysql.Open(configs.GetTestDBInfo()), &gorm.Config{
		Logger:                 newLogger,
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
}

func InitTestRedis() {
	redis.Init()
}

func TestInit() {
	InitTestDb()
	InitTestRedis()
	dao.DB = testDb
}
