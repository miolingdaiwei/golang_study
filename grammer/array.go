package main

import "fmt"

func main() {
	var a [3]int // 定义三个整数的数组
	//不赋值的话就默认值为0
	fmt.Println(a[0])        // 打印第一个元素
	fmt.Println(a[len(a)-1]) // 打印最后一个元素
	// 打印索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	// 仅打印元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var r = [3]int{1, 2}
	fmt.Println(r[2]) // "0"

	sl := r[0:1]
	fmt.Println("slice is ", sl)
	//var lessa = [3]int{1, 2, 3}
	lessa := [...]int{1, 2, 3, 4}
	fmt.Println(lessa)

	//遍历数组
	var arr = [3]int{2, 3, 5}
	for k, v := range arr {
		fmt.Println(k, v)
	}
	//two-dimensional array
	var doubleArr = [...][2]int{{3, 3}, {2, 2}, {2, 4}}
	fmt.Println(doubleArr)
	for k, v := range doubleArr {

		for k, v := range v {
			fmt.Print(k, v, " ")
		}
		fmt.Println(k, v)
	}

	// 声明两个二维整型数组
	var array1 [2][2]int
	var array2 [2][2]int
	// 为array2的每个元素赋值
	array2[0][0] = 10
	array2[0][1] = 20
	array2[1][0] = 30
	array2[1][1] = 40
	// 将 array2 的值复制给 array1
	array1 = array2

	fmt.Println(array1, array2)
	//	深拷贝

	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}
}
