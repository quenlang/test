package main

import (
	"fmt"
	"os"
)

/*  这种变量声明初始化的方式只能用于func内部，全局变量的声明和初始化不能使用这种方式
os := "windows"
*/

// 全局变量即便后边没有用到也允许声明，但局部变量则不允许，会编译报错
var a, b, c int = 0, 1, 2

func main() {
	var isDown bool = true
	var age = 26
	gopath := os.Getenv("GOPATH")

	fmt.Printf("isDown: %v, age: %d, gopath: %s\n", isDown, age, gopath)
}
