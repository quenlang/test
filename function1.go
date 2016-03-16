package main

import (
	"fmt"
)

/*
建议使用命名返回值， 可读性更高
*/

func main() {
	a, b, c := oprate(10, 2)
	fmt.Printf("+ : %d\n", a)
	fmt.Printf("- : %d\n", b)
	fmt.Printf("* : %d\n", c)
}

/*
func oprate(a, b int) (m, n, k int) {
    m = a + b
    n = a - b
    k = a * b

    return
}
*/

func oprate(a, b int) (int, int, int) {
	return a + b, a - b, a * b
}
