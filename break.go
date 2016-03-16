package main

func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if j > 5 {
				//break 只能退出最内层的循环
				break
			}
			print(j)
		}
		print(" ")
	}
}
