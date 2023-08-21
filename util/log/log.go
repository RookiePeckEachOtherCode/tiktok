package tiktokLog

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Normal 普通log
func Normal(format string, values ...any) {
	now := time.Now().Format("2006/01/02 - 15:04:05")
	f := fmt.Sprintf("[Normal] %s %s\n", now, format)
	fmt.Fprintf(gin.DefaultWriter, f, values...)
}

// Error 错误log
func Error(ErrorInfo string, values ...any) {
	now := time.Now().Format("2006/01/02 - 15:04:05")
	f := fmt.Sprintf("[Error] %s %s %v \n", now, ErrorInfo, values)
	fmt.Fprintf(gin.DefaultWriter, f, values...)
}
