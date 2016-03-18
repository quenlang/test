package main

import (
	"fmt"
)

/*
函数作为其他函数的参数
*/

func main() {
	callback(1, Add)
}

func Add(m, n int) {
	fmt.Printf("The sum of %d and %d is: %d\n", m, n, m+n)
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}
