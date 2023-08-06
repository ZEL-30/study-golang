package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat() ...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk() ...")
}

type SuperMan struct {
	Human // SuperMan 类继承了Human类的方法

	level int
}

// 重定义父类方法
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat() ...")
}

// 子类新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly() ...")
}

func main() {

	human := Human{"赵云", "男"}
	human.Eat()
	human.Walk()

	// 子类实例化
	// superman := SuperMan{Human{"钢铁侠", "男"}, 100}
	var superman SuperMan
	superman.name = "钢铁侠"
	superman.sex = "男"
	superman.level = 100
	superman.Eat()
	superman.Fly()
	fmt.Println(superman)
}
