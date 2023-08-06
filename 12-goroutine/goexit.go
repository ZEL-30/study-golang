package main

import (
	"fmt"
	"runtime"
)

func main() {

	// 用 go 创建承载一个形参为空, 返回值为空的一个函数
	go func() {
		defer fmt.Println("B.defer")
		runtime.Goexit()
		fmt.Println("B")
	}()

	fmt.Println("A")

	for {
	}

}
