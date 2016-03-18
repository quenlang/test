package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	f := Adder()
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	fmt.Println(f(1))
	fmt.Println(f(20))
	where()
	fmt.Println(f(300))
}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
