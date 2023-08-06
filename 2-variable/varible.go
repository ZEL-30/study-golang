// 四种变量的声明方式

package main

import "fmt"

func main() {

	var a int
	fmt.Println(a)

	var b int = 100
	fmt.Println(b)

	var c = 100
	fmt.Printf("Type of c is %T\n", c)

	// 省略var关键字， 直接自动匹配
	d := 300
	fmt.Println(d)

	var x, y, z int = 10, 20, 40
	fmt.Printf("x = %d, y = %d, z = %d\n", x, y, z)

}
