package main

/*
实际中并不建议使用标签和goto的组合
*/

func main() {
	i := 0
HERE:
	print(i)
	i++
	if i == 5 {
		return
	}

	goto HERE
}
