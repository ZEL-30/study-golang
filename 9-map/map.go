package main

import "fmt"

func printMap(map1 map[int]string) {
	// map1 是一个引用传递
	for key, value := range map1 {
		fmt.Println("key = ", key, "value = ", value)
	}
}

func main() {

	// 声明map1是一种map类型， key是string, value是string
	map1 := map[string]string{
		"one":   "C++",
		"two":   "Java",
		"three": "Golang",
	}
	fmt.Println(map1)

	// 添加
	map2 := make(map[int]string)
	map2[1] = "C++"
	map2[2] = "Golang"
	map2[3] = "Python"
	fmt.Println(map2)
	fmt.Println("----------------------")

	// 删除
	delete(map2, 2)

	// 遍历
	for key, value := range map2 {
		fmt.Println("key = ", key, "value = ", value)
	}

}
