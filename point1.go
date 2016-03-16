package main

import (
	"fmt"
)

func main() {
	var i1 int = 5
	fmt.Printf("An integer: %d,its location in memeory: %p\n", i1, &i1)
	var ptr *int
	ptr = &i1
	fmt.Printf("The value at memory location %p is %d\n", ptr, *ptr)
}
