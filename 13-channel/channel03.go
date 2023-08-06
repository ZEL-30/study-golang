package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		// close 可以关闭一个 channel
		close(c)
	}()

	// for {
	// 	// ok: 如果为 true 表示 channel 没有关闭, 如果为 false 表示 channel 已经关闭
	// 	if data, ok := <-c; ok {
	// 		fmt.Println(data)
	// 	} else {
	// 		break
	// 	}
	// }

	// 可以使用 range 来迭代不断操作 channel
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Finished..")
}
