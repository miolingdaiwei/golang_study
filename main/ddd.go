package main

import "fmt"

func ddd(int1 ...int) {
	for _, i2 := range int1 {
		fmt.Println(i2)
	}
}

func ddd2(int2 ...int) {
	ddd(int2...)
	//通过这个传回可变参数
}

func main() {
	ddd2(1, 2, 3, 4)
}
