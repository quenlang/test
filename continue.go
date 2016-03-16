package main

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			//忽略本次循环，进入下一次循环，该关键字只能用于for循环中
			continue
		}
		print(i)
		print(" ")
	}
}
