package logprint

import (
	"fmt"
	"time"
)

func Info(msg interface{}) {
	t := time.Now()
	// 定义info日志格式
	fmt.Printf("[info] %s: %s\n", t.Format("2006-01-02 15:04:05"), msg)
}
