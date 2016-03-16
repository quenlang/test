package main

import (
	"fmt"
)

func main() {
	x := min(1, 3, 4, 2, 0)
	fmt.Printf("the mininum is %d\n", x)
	arr := []int{7, 9, 3, 5, 1}
	// 如果参数被存储在一个数组 arr 中，则可以通过 arr... 的形式来传递参数调用变参函数
	x = min(arr...)
	fmt.Printf("the mininum is %d\n", x)
}

func min(a ...int) (min int) {
	if len(a) == 0 {
		min = 0
	}
	min = a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return
}
