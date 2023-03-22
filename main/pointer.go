package main

import (
	"flag"
	"fmt"
)

//指针（pointer）在Go语言中可以被拆分为两个核心概念：
//类型指针，允许对这个指针类型的数据进行修改，传递数据可以直接使用指针，而无须拷贝数据，类型指针不能进行偏移和运算。
//切片，由指向起始元素的原始指针、元素数量和容量组成。

//C/C++ 中指针饱受诟病的根本原因是指针的运算和内存释放，C/C++ 语言中的裸指针可以自由偏移，甚至可以在某些情况下偏移进入操作系统的核心区域，
//我们的计算机操作系统经常需要更新、修复漏洞的本质，就是为解决指针越界访问所导致的“缓冲区溢出”的问题。

func swapTrue(a, b *int) {
	// 取a指针的值, 赋给临时变量t
	t := *a
	// 取b指针的值, 赋给a指针指向的变量
	*a = *b
	// 将a指针的值赋给b指针指向的变量
	*b = t
}

func swapFalse(a, b *int) {
	b, a = a, b
}

var mode = flag.String("mode", "", "process mode")

func main() {

	// 准备一个字符串类型
	var house = "Malibu Point 10880, 90265"
	// 对字符串取地址, ptr类型为*string  *string就是指针的类型
	ptr := &house
	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)
	// 打印ptr的指针地址
	fmt.Printf("address: %p\n", ptr)
	// 对指针进行取值操作
	value := *ptr
	// 取值后的类型
	fmt.Printf("value type: %T\n", value)
	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)

	flag.Parse()
	// 输出命令行参数
	fmt.Println(*mode)

	str := new(string)
	*str = "Go语言教程"
	fmt.Println(*str)
}
