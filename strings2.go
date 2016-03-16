package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "oink oink oink"
	str1 := " Amoeba Cobar "
	fmt.Println(strings.Replace(str, "k", "ky", 2))
	fmt.Println(strings.Count(str, "o"))
	fmt.Println(strings.Count(str, ""))
	fmt.Println(strings.Count(str, " "))
	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.ToLower(str1))
	fmt.Println(str1)
	fmt.Println(strings.TrimSpace(str1))

}
