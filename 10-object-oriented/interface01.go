package main

import "fmt"

// 本质是一个指针
type IAnimal interface {
	Sleep()
	GetColor() string // 获取动物的颜色
	GetType() string  // 获取动物的种类
}

// 具体的类
type Cat struct {
	color string // 猫的颜色
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string // 猫的颜色
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func main() {

	var animal IAnimal
	animal = &Cat{"黄色"}
	animal.Sleep()
	fmt.Println(animal.GetColor())
}
