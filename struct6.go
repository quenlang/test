package main

import (
	"fmt"
)

type B struct {
	a int
	b int
}

type D struct {
	B
	b float32
}

func main() {
	d := D{B{1, 3}, 14.2}
	fmt.Println(d.b)
	fmt.Println(d.a)
	fmt.Println(d.B.a)
	fmt.Println(d.B.b)
}
