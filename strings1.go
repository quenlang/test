package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "go is a amazing language !"
	hasPre := strings.HasPrefix(str, "go")
	hasSuf := strings.HasSuffix(str, "!")
	isContain := strings.Contains(str, "amazing")
	index := strings.Index(str, "l")
	lastIdx := strings.LastIndex(str, "a")
	idxRune := strings.IndexRune(str, 'a')

	fmt.Printf("hasPre`s value: %v\n", hasPre)
	fmt.Printf("hasPre`s value: %v\n", hasSuf)
	fmt.Printf("hasPre`s value: %v\n", isContain)
	fmt.Printf("hasPre`s value: %v\n", index)
	fmt.Printf("hasPre`s value: %v\n", lastIdx)
	fmt.Printf("hasPre`s value: %v\n", idxRune)
}
