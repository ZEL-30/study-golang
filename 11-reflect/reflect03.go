package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (user User) Call() {
	fmt.Println("Call is called ...")
	fmt.Printf("%v\n", user)
}

func (user User) Show() {
	fmt.Println("Show is called ...")
	fmt.Printf("%v\n", user)
}

func main() {

	user := User{1, "zel", 25}

	DoFieldAndMethod(user)

}

func DoFieldAndMethod(input interface{}) {
	// 获取 input 的 Type 和 Value
	input_type := reflect.TypeOf(input)
	input_value := reflect.ValueOf(input)

	// 通过 type 获取里面的字段
	// 1. 获取interface的reflect.Type, 通过Type得到numField, 进行遍历
	// 2. 得到每个field, 数据类型
	// 3. 通过field有一个Interface()方法得到对应的value
	for i := 0; i < input_type.NumField(); i++ {
		field := input_type.Field(i)
		value := input_value.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 通过 type 获取里面的方法并调用
	fmt.Println(input_type.NumMethod())
	for i := 0; i < input_type.NumMethod(); i++ {
		method := input_type.Method(i)
		fmt.Printf("%s: %v\n", method.Name, method.Type)
	}

}
