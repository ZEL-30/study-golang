package main

import "fmt"

// 声明一种新的数据类型，myint, 是int的一个别名
type myint int

// 定义一个结构体
type Book struct {
	title string
	auth  string
}

func printBook(book Book) {
	fmt.Printf("书名: %s, 作者: %s\n", book.title, book.auth)
}

func changeBook(book *Book) {
	// 结构体转递是值传递
	book.title = "活着"
	book.auth = "余华"
}

func main() {

	var book Book
	book.title = "三体"
	book.auth = "刘慈欣"

	changeBook(&book)

	printBook(book)
}
