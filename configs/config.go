package configs

import "fmt"

const SIGN_KEY = "RookiePeckEachOtherCode"

const MAX_VIDEO_CNT = 30

// mysql的连接信息
const (
	DB_USER   = "root"
	DB_PASSWD = "db22455"
	DB_URL    = "127.0.0.1"
	PORT      = "3306"
	DB_NAME   = "tiktok"
)

func GetDBInfo() string {
	return DB_USER + ":" + DB_PASSWD + "@tcp(" + DB_URL + ":" + PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
}

// redis的连接信息
const (
	RDB_IP   = "127.0.0.1"
	RDB_PORT = "6379"
)

func GetRedisInit() string {
	return fmt.Sprintf("%s:%s", RDB_IP, RDB_PORT)
}
