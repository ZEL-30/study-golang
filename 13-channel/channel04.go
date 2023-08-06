package main

import "fmt"

func fibonacii(c, quit chan int) {
	x, y, temp := 1, 1, 0

	for {
		select {

		case c <- x:
			// 如果 c 可写, 则该 case 就会进来
			temp = x
			x = y
			y += temp

		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {

	c := make(chan int)

	quit := make(chan int)

	// sub go
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	// main go
	fibonacii(c, quit)
}
