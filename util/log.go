package util

import (
	"fmt"
	"strings"
)

func PrintLog(log string) {
	// 计算日志信息的长度
	length := len(log) + 4

	// 打印上边框
	fmt.Println(strings.Repeat("+", length))

	// 打印日志信息
	fmt.Printf("| %s |\n", log)

	// 打印下边框
	fmt.Println(strings.Repeat("+", length))
}
