package main

//import (
//	"bytes"
//	"fmt"
//	"io"
//	"os"
//)
//
//func main() {
//	var x interface{}
//	x = 10
//	//断言能够得到value和Boolean的返回值
//	value, ok := x.(int)
//	fmt.Print(value, ",", ok, "\n")
//
//	var y interface{}
//
//	y = "string"
//	//通过变量.(type)来switch...case
//	switch y.(type) {
//	case int:
//		fmt.Println("int")
//	case string:
//		fmt.Println("string")
//	case float32:
//		fmt.Println("float32")
//	default:
//		fmt.Println("?")
//	}
//
//	var w io.Writer
//	w = os.Stdout
//	f := w.(*os.File)
//	//t := w.(*bytes.Buffer)
//	//只读取一个参数的话如果w没有实现断言的类型的接口，那么就会panic
//	t2, ok := w.(*bytes.Buffer)
//	//两个参数时，ok会读取到w有没有实现接口的bool值，不会panic
//	fmt.Println(f, t2)
//}

import "fmt"

// Flyer 定义飞行动物接口
type Flyer interface {
	Fly()
}

// Walker 定义行走动物接口
type Walker interface {
	Walk()
}

// bird 定义鸟类
type bird struct {
}

// Fly 实现飞行动物接口
func (b *bird) Fly() {
	fmt.Println("bird: fly")
}

// Walk 为鸟添加Walk()方法, 实现行走动物接口
func (b *bird) Walk() {
	fmt.Println("bird: walk")
}

// 定义猪
type pig struct {
}

// Walk 为猪添加Walk()方法, 实现行走动物接口
func (p *pig) Walk() {
	fmt.Println("pig: walk")
}
func main() {

	//tt := map[string]int{
	//	"dsds": 2,
	//}
	// 创建动物的名字到实例的映射
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}
	// 遍历映射
	for name, obj := range animals {
		// 判断对象是否为飞行动物
		f, isFlyer := obj.(Flyer)
		// 判断对象是否为行走动物
		w, isWalker := obj.(Walker)
		fmt.Printf("name: %s isFlyer: %v isWalker: %v\n", name, isFlyer, isWalker)
		// 如果是飞行动物则调用飞行动物接口
		if isFlyer {
			f.Fly()
		}
		// 如果是行走动物则调用行走动物接口
		if isWalker {
			w.Walk()
		}
	}
}
