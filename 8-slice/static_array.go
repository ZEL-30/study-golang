package main

import "fmt"

func printArray(array [10]int) {

	for index, value := range array {
		fmt.Println("index = ", index, "value = ", value)
	}

}

func main() {

	var array1 [10]int

	for i := 0; i < len(array1); i++ {
		array1[i] = 10
	}

	fmt.Println(array1)

	array2 := [10]int{1, 2, 3, 4}
	printArray(array2)

	// 查看数组的数据类型
	fmt.Printf("array1 = %d, Type of array1 is %T\n", array1, array2)

}
