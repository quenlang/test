package main

import (
	"fmt"
)

func main() {
	var arr = []int{0, 1, 2, 3, 4, 5}
	//数组切片化
	s := sum(arr[:])
	fmt.Printf("Result is :%d\n", s)
}

func sum(a []int) (res int) {
	for i := 0; i < len(a); i++ {
		res += a[i]
	}
	return
}
