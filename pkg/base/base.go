package base

import (
	"fmt"
	"runtime"
)

func PrintCallFunc() {
	// 获取调用者的信息
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("无法获取调用者信息")
		return
	}

	// 获取函数名
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		fmt.Println("无法获取函数名")
		return
	}

	fmt.Printf("当前执行位置: %s\n", fn.Name())
	fmt.Printf("文件: %s, 行号: %d\n", file, line)
}
