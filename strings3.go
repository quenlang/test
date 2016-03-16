package main

import (
	"fmt"
	"strings"
)

/*
strings.Fields(s)
strings.Split(s, sep)
strings.Join(a, sep)
*/
func main() {
	str := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()

	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)

	fmt.Println(strings.Join(sl, ","))
	fmt.Println(strings.Join(sl2, "-"))

}
