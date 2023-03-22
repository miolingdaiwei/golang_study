package main

import "fmt"

// 匿名函数作为返回值
func Add(i int) func() int {
	return func() int {
		i++
		return i
	}
}

func main() {
	//获取到返回的函数
	a1 := Add(10)

	//然后调用返回的函数，得到操作之后的结果，这就是go的函数实现的闭包
	fmt.Println(a1()) // 13
}
