package main

import "fmt"

// interface{} 是万能数据类型
func myFunc(arg interface{}) {
	// interface{} 改如何区分， 此时引用的底层数据类型到底是什么
	// 给 interface{} 提供 "类型断言" 的机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string")
	} else {
		fmt.Println("arg is string type, value = ", value)
		fmt.Printf("value type is %T\n", value)
	}

}

func main() {

	myFunc(100)
	myFunc("ZEL")

}
