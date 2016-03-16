package main

import (
	"fmt"
	"runtime"
)

var prompt string = "Enter a digit, e.g. 3 " + "or %s to quit."

func init() {
	if "windows" == runtime.GOOS {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else {
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

func main() {
	fmt.Println(prompt)
}
