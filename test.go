package main

import (
	"fmt"
)

var a uint32

func main() {
	a = 1 << 20
	fmt.Print("Value: %d\n", a)
}
