package dao

import (
	"log"

	"tiktok/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {
	var err error
	dsn := configs.GetDBInfo()

	DB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Panicln("err :", err.Error())
	}

}
