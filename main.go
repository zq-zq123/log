package main

import (
	"fmt"
	"logproject/logprint"
)

func main() {
	fmt.Println("程序开始启动...")
	fmt.Println("运行中...")
	fmt.Println("程序运行结束")
	// 调用自定义包的方法，进行打印日志
	logprint.Debug("这是一个debug日志")
	logprint.Info("这是一个info日志")
	logprint.Warn("这是一个warn日志")
}
