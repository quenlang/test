package main

import (
	"fmt"
)

type A struct {
	ax, ay int
}

type B struct {
	A
	bx, by int
}

func main() {
	// 该初始化方法：嵌套结构体必须指明类型
	b := B{A{1, 2}, 3.0, 4.0}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)
}
