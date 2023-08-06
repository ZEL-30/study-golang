package main

import "fmt"

func main() {

	numbers := make([]int, 3, 10)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向numbers 切片追加一个元素1, numbers len = 4, [0, 0, 0, 1], cap = 10
	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 深拷贝
	numbersCopy := make([]int, len(numbers))
	copy(numbersCopy, numbers)

	// 切片的截取
	s1 := numbersCopy[0:2]
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(s1), cap(s1), s1)

}
