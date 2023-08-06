package main

import "fmt"

// const 定义枚举类型
const (
	// 可以在const() 添加一个关键字 iota, 每行的iota都会累加1, 第一行的iota默认值是0
	// iota 只能配合const 进行使用
	BEIJING = iota + 1
	SHANGHAI
	SHENZHEN
)

func main() {

	const length = 10

	fmt.Println(length)
	fmt.Println("北京 =", BEIJING)
	fmt.Println("上海 =", SHANGHAI)
	fmt.Println("深圳 =", SHENZHEN)

}
