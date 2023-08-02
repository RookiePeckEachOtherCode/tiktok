package configs

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
