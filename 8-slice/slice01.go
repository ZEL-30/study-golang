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
	// 声明slice1是一个切片, 并且初始化，默认值是1,2,3, 长度len是3
	slice1 := []int{1, 2, 3}
	fmt.Printf("len = %d, slice1 = %v\n", len(slice1), slice1)

	var slice2 []int
	slice2 = make([]int, 2)
	slice2[0] = 100
	fmt.Printf("len = %d, slice2 = %v\n", len(slice2), slice2)

	slice3 := make([]int, 100)
	fmt.Printf("len = %d, slice3 = %v\n", len(slice3), slice3)

	// 判断一个slice是否为0
	if slice3 == nil {
		fmt.Println("slice3 是一个空切片")
	} else {
		fmt.Println("slice3 是有空间的")
	}

}
