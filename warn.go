// 包名最好和上级目录保持一致，最好不要带下划线
package logprint

import (
	"fmt"
	"time"
)

func Warn(msg interface{}) {
	t := time.Now()
	// 定义warn日志格式
	fmt.Printf("[warn] %s: %s\n", t.Format("2006-01-02 15:04:05"), msg)
}
