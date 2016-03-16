package main

import (
	"fmt"
)

/*
如果要用goto就用正序，需要注意的是goto和LABEL之间不能有新的变量声明
*/
func main() {
	a := 1
	var b int
	goto TARGET
	//b := 9
TARGET:
	b += a
	fmt.Printf("a is %v *** b is %v", a, b)
}
