package pid

import (
	"os"
	"strconv"
)

// pid文件名称
var filePath = "pid"

// 保存pid
func init() {
	err := os.WriteFile(filePath, []byte(strconv.Itoa(os.Getpid())), 0666)
	if err != nil {
		panic(err)
	}
}
