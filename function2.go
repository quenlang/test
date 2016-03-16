package main

import (
	"fmt"
)

/*
利用指针变量参数修改函数外部的变量
*/

func main() {
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Printf("reply`s value: %v\n", *reply)
}

func Multiply(a, b int, reply *int) {
	*reply = a * b
}
