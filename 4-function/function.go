package main

import "fmt"

func fool1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100

	return c
}

// 返回多个返回值， 匿名的
func fool2(a string, b int) (int, string) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100
	d := "zel-30"

	return c, d
}

// 返回多个返回值， 命名的
func fool3(a string, b int) (c int, d string) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c = 100000000000
	d = "zel-30"

	return c, d
}

func main() {

	result_1, result_2 := fool2("zel", 30)
	fmt.Println(result_1, result_2)

	result_1, result_2 = fool3("zel", 30)
	fmt.Println(result_1, result_2)

}
