//package main
//
//import "fmt"
//
//func main() {
//	var array1 = [4]int{2, 3, 4}
//	s1 := array1[1:]
//	s2 := array1[0:1]
//	println(array1, s1, s2, array1[:])
//	var highRiseBuilding [30]int
//	for i := 0; i < 30; i++ {
//		highRiseBuilding[i] = i + 1
//	}
//	// 区间
//	fmt.Println(highRiseBuilding[10:15])
//	// 中间到尾部的所有元素
//	fmt.Println(highRiseBuilding[20:])
//	// 开头到中间指定位置的所有元素
//	fmt.Println(highRiseBuilding[:2])
//}

package main

import "fmt"

func main() {
	var arr1 = [3]int{2, 3, 4}
	fmt.Println(arr1[1:], arr1[:], arr1[1:2], arr1[1:1])
	//[3 4] [2 3 4] [3] []
	var slice1 []int
	println(len(slice1))
	s2 := make([]int, 3)
	//s2[0] = 1
	//s2[1] = 1
	//s2[2] = 2
	s2_ := append(s2, 3)
	s2_ = append(s2_, 4)
	fmt.Println(s2, s2_, len(s2_))
}
