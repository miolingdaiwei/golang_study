package main

import (
	"fmt"
	"time"
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {

	const (
		a = 1
		b
		c = 2
		d
	)
	//const 批量声明时，除了第一个变量需要赋值，后面都可以省略

	const c1 = 2 / 3 //正确的做法
	//const c2 = getNumber()   错误的做法 引发构建错误: getNumber() 用做值
	//const 变量的值需要能在编译阶段确定

	fmt.Print(Wednesday)
	fmt.Print(Tuesday)

	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s"
	fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"
	fmt.Printf("%T %v\n", 1, "b")
	fmt.Printf("%T %[1]v\n", 1, "s")
	fmt.Printf("%T %v", "b", 2)
	//%T 打印参数类型
	fmt.Print(a, b, c, d)
}
