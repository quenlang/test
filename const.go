package main

import (
	"fmt"
)

/*
iota 总是伴随这const出现的，且每出现一次值就加一，再次遇到const关键字被重置为0
*/
func main() {

	const (
		a = iota
		b
		c
	)
	/*等同于如下格式
	  const (
	      a = iota
	      b = iota
	      c = iota
	  )
	*/
	const (
		d = iota
	)

	fmt.Printf("a: %d, b: %d, c: %d, d: %d\n", a, b, c, d)
}
