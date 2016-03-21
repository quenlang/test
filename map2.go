package main

import (
	"fmt"
)

func main() {
	items := make([]map[int]int, 5)
	for i := range items {
		fmt.Println(items[i])
		items[i] = make(map[int]int, 2)
		items[i][1] = 2
		items[i][2] = 4
	}

	fmt.Printf("Version A: Value of items: %v\n", items)
}
