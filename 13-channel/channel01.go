package main

import (
	"fmt"
)

// 无缓冲的 channel

func main() {

	// 定义一个channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine 结束运行")
		fmt.Println("goroutine 正在运行 ...")

		c <- 666 // 将 666 发送到 c
	}()

	num := <-c // 从 c 中接受数据， 并赋值给num

	fmt.Println("num =", num)
	fmt.Println("main goroutine 结束运行")

}
