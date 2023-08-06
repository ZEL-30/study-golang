package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Println("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {

	// 创建一个goroutine 去执行newTask() 流程
	go newTask()

	time.Sleep(10 * time.Second)
	fmt.Println("main exit")

}
