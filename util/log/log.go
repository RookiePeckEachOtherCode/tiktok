package tiktokLog

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var now string

// Log 普通log
func Normal(format string, values ...any) {
	now = time.Now().Format("2006/01/02 - 15:04:05")
	f := fmt.Sprintf("[DEV] %s %s\n", now, format)
	fmt.Fprintf(gin.DefaultWriter, f, values...)
}

// LogError 带错误信息的log(服务器的错误)
func Error(ErrorInfo string, values ...any) {
	now = time.Now().Format("2006/01/02 - 15:04:05")
	f := fmt.Sprintf("[DEV] [Error] %s %s %v \n", now, ErrorInfo, values)
	fmt.Fprintf(gin.DefaultWriter, f)
}

// LogFatal 严重的错误
func Fatal(ErrorInfo string, values ...interface{}) {
	now = time.Now().Format("2006/01/02 - 15:04:05")
	f := fmt.Sprintf("[DEV] [Fatal] %s %s %v\n", now, ErrorInfo, values)
	fmt.Fprint(gin.DefaultWriter, f)
}
