package main

import (
	"fmt"
	"net"
)

func GetData() (int, int) {
	return 100, 200
}

func main() {
	//var name type
	var variate string
	//多行声明
	variate = "vari"
	var (
		a int
		b string
		c []float32
		//d func() bool
		e struct {
			x int
		}
	)
	//名字 := 表达式
	f := "ss"
	i, j := "1", "c"

	//f:="dd"  打咩，不能重复声明，底层就是变量名单一

	//使用短变量声明，且有新变量在左边，就不会报错
	//大概短变量声明的左边的多个变量名是存在同一个新的区域的吧！
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	conn2, err := net.Dial("tcp", "127.0.0.1:8080")

	//通过多变量赋值来进行变量交换  因为早go里面内存空间不再是紧缺的资源。
	i, j = j, i

	//func Printf(format string, a ...any) (n int, err error)
	fmt.Printf("hello" + "   " + variate)
	//d()
	fmt.Print(a, b, c, e, f, i, j, conn, conn2, err)

	m, _ := GetData()
	_, n := GetData()
	fmt.Print(m, n)
	//100 200   _是匿名变量，不使用内存
	s := ""
	fmt.Print(s != "" && s[0] == 'x')
	//左边为真，那么&&虽然s[0]不能用于空字符串，但是也可以用
}
