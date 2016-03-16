package main

import (
	"fmt"
)

func main() {
	s := "good bye"
	var p *string = &s
	// 反向引用;注意：对于一个空指针的反向引用是不合法的。runtime error: invalid memory address or nil pointer dereference
	*p = "ciao"

	fmt.Printf("Here is the pointer p: %p\n", p)
	fmt.Printf("Here is the string *p: %v\n", *p)
	fmt.Printf("Here is the string s: %s\n", s)
}
