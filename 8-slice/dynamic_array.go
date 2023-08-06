package main

import "fmt"

func printArray(array *[]int) {
	// _ 表示匿名的变量
	for _, value := range *array {
		fmt.Println("value = ", value)
	}

	(*array)[1] = 10
}

func main() {
	array := []int{1, 2, 3, 4}
	printArray(&array)
	fmt.Println(array)

}
