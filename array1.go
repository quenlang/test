package main

import (
	"fmt"
)

/*
数组赋值需要做一次内存拷贝
*/

func main() {
	var arr1 = [5]int{1, 2, 3, 4, 5}
	//arr2 := &arr1
	arr2 := arr1
	arr2[1] = 200
	fmt.Printf("arr1[1] value: %d\n", arr1[1])
	fmt.Printf("arr2[1] value: %d\n", arr2[1])
}
