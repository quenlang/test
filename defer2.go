package main

import (
	"fmt"
)

func main() {
	b()
}

/*
调试过程中做代码追踪
*/
func trace(s string) {
	fmt.Println("entering: ", s)
}

func untrace(s string) {
	fmt.Println("leaving: ", s)
}

func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a()
}
