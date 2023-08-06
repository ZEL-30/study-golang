package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {

	a, b := 10, 20
	fmt.Println("a = ", a, "b = ", b)

	swap(&a, &b)
	fmt.Println("a = ", a, "b = ", b)

}
