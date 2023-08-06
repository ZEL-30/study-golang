package main

import (
	_ "golang-study/5-init/lib-1"    // 给 lib1 包起一个别名，匿名， 无法使用当前包的方法， 但是会执行当前的包内部的init()方法
	lib2 "golang-study/5-init/lib-2" // 给 lib2 包起一个别名 lib2, 用 lib2.Lib2Test(), 来调用
)

func main() {

	lib2.Lib2Test()
}
