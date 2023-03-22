package main

import "fmt"

// a func with a func parameter
func userCallBackFunc(i int, f func(int2 int)) {
	f(i)
}

func main() {
	//在这里传入匿名函数
	userCallBackFunc(3, func(int2 int) {
		fmt.Printf("do callBack and the number is  %d  \n", int2)
	})
}
