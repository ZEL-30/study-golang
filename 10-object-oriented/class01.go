package main

import "fmt"

// 如果类名首字母大写， 表示其他包也能够访问
type Hero struct {
	// 如果说类的属性首字母大写， 表示该属性是对外能够访问的， 否则的话只能够类的内部访问
	name  string
	ad    int
	level int
}

func (this *Hero) GetName() string {
	return this.name
}

func (this *Hero) SetName(name string) {
	this.name = name
}

func main() {

	// 创建一个对象
	hero := Hero{name: "赵云", ad: 100, level: 1}
	fmt.Println("名字: ", hero.GetName())

	hero.SetName("张恩乐")
	fmt.Println(hero)

}
