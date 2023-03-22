package main

import "fmt"

// go也注重类型，接口就是强化类型的重要性

//type 接口类型名 interface{
//	方法名1( 参数列表1 ) 返回值列表1
//	方法名2( 参数列表2 ) 返回值列表2
//}

// Word 定义一个接口，有outWord方法
type Word interface {
	outWord(data string) string
}

// Words 定义一个结构，他将要实现接口
type Words struct {
}

//定义一个方法，它和interface里的outWord定义的函数一致 （注意：这是一个方法而不是函数，words作为接收着）
func (w *Words) outWord(data string) string {
	fmt.Println("outWords: ", data)
	return " s "
}

type add interface {
	addOne() bool
}

type del interface {
	delOne() bool
}

type num interface {
	add
	del
}

type number struct {
}

func (n *number) add() bool {
	return true
}

func (n *number) del() bool {
	return true
}

func main() {
	//定义好一个结构体变量，使用new分配空间
	w := new(Words)

	//定义一个接口变量
	var word Word

	//结构体实现接口，能够赋值的前提是w （words）实现了interface里定义的成员 （注意：必须全部实现才可以）
	word = w

	//那么这个被实现后的接口变量就可以调用outWord方法
	word.outWord("data")

	var x interface{}
	x = 1
	//var y int =x
	//上面这种写法是不可以的，因为x的类型还是interface
	var y, ok = x.(int)
	//可以通过断言拿到值,但是断言的是int,且bool是true
	fmt.Println(y, ok) //1 true

}
