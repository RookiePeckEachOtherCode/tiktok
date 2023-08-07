package configs

import (
	"fmt"
)

// gin的ip和端口
const GIN_IP = "127.0.0.1"
const GIN_PORT = 8080

// SIGN_KEY 是用于签名的密钥
const SIGN_KEY = "RookiePeckEachOtherCode"

// MAX_VIDEO_CNT 是每个用户最多可以上传的视频数量
const MAX_VIDEO_CNT = 30

// mysql的连接信息
const (
	DB_USER   = "root"      // 数据库用户名
	DB_PASSWD = "db22455"   // 数据库密码
	DB_URL    = "127.0.0.1" // 数据库地址
	PORT      = "3306"      // 数据库端口
	DB_NAME   = "tiktok"    // 数据库名称
)

// GetDBInfo 返回mysql的连接信息
func GetDBInfo() string {
	return DB_USER + ":" + DB_PASSWD + "@tcp(" + DB_URL + ":" + PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
}

// redis的连接信息
const (
	RDB_IP   = "127.0.0.1" // redis地址
	RDB_PORT = "6379"      // redis端口
)

// GetRedisInit 返回redis的连接信息
func GetRedisInit() string {
	return fmt.Sprintf("%s:%s", RDB_IP, RDB_PORT)
}

// 注册用户名，最长32个字符
const MAX_NAME_LEN = 32

// 密码，最长32个字符
const MAX_PASSWORD_LEN = 32

// Jwt密钥
const JWT_KEY = "RookiePeckEachOtherCode"

// VideoSavePath 视频保存路径
const VIDEO_SAVE_PATH = "/static/assets/video/"

// VideoCoverSavePath 视频封面保存路径
const VIDEO_COVER_SAVE_PATH = "/static/assets/video_cover/"

func Bless() {
	fmt.Println("                    _ooOoo_")
	fmt.Println("                   o8888888o")
	fmt.Println("                   88\" . \"88")
	fmt.Println("                   (| -_- |)")
	fmt.Println("                    O\\ = /O")
	fmt.Println("                ____/`---'\\____")
	fmt.Println("              .   ' \\| |// `.")
	fmt.Println("               / \\||| : |||// \\")
	fmt.Println("             / _||||| -:- |||||- \\")
	fmt.Println("               | | \\\\\\ - /// | |")
	fmt.Println("             | \\_| ''\\---/'' | |")
	fmt.Println("              \\ .-\\__ `-` ___/-. /")
	fmt.Println("           ___`. .' /--.--\\ `. . __")
	fmt.Println("        .\"\" '< `.___\\_<|>_/___.' >'\"\".")
	fmt.Println("       | | : `- \\`.;`\\ _ /`;.`/ - ` : | |")
	fmt.Println("         \\ \\ `-. \\_ __\\ /__ _/ .-` / /")
	fmt.Println(" ======`-.____`-.___\\_____/___.-`____.-'======")
	fmt.Println("                    `=---='")
	fmt.Println("")
	fmt.Println(".............................................")
	fmt.Println("          佛祖保佑             永无BUG")
}
