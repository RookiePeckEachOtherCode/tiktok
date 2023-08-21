package dao_test

import (
	"io"
	"log"
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
		log.New(io.Discard, "\r\n", log.LstdFlags), // 不输出日志
		logger.Config{
			SlowThreshold: time.Second,   // 慢SQL阈值
			LogLevel:      logger.Silent, // 日志级别
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
	redis.RedisFlushAll()
}

func TestInit() {
	InitTestDb()
	InitTestRedis()
	dao.DB = testDb
}
