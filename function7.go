package main

import (
	"fmt"
)

/*
  0- g is of type func(int) and has value 0x401220
  1- g is of type func(int) and has value 0x401220
  2- g is of type func(int) and has value 0x401220
  3- g is of type func(int) and has value 0x401220
*/

func main() {
	f()
}

func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) {
			fmt.Printf("%d", i)
		}
		g(i)
		fmt.Printf("- g is of type %T and has value %v\n", g, g)
	}
}
