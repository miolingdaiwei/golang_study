//package main
//
//import "fmt"
//
//type People struct {
//	name  string
//	child *People
//	//	People结构体的指针
//}
//
//func tost(p1 *People) {
//	fmt.Println(p1)
//}
//
//func main() {
//	type Name struct {
//		firstName string
//		lastName  string
//	}
//	type Person struct {
//		name Name
//		age  int
//	}
//	var person1 Person
//	person1.age = 19
//	person1.name.firstName = "m"
//	person1.name.lastName = "io"
//
//	fmt.Println(person1, person1.name.firstName+person1.name.lastName)
//	//{{m io} 19} mio
//
//	//type People struct {
//	//	name  string
//	//	child *People
//	//	//	People结构体的指针
//	//}
//	//relation := &People{
//	//	//取到这个结构体实例的地址
//	//	name: "爷爷",
//	//	child: &People{
//	//		name: "爸爸",
//	//		child: &People{
//	//			name: "我",
//	//		},
//	//	},
//	//}
//
//	//fmt.Println(relation, *relation.child, *(*relation.child).child)
//	////&{爷爷 0xc000004090} {爸爸 0xc0000040a8} {我 <nil>}
//	////	值肯定是地址，但是地址指向的区域包括name和另一个地址
//	var p1 People
//	p1.name = "ss"
//	tost(&p1)
//}
package main

import "fmt"

type Str struct {
	a int
	b int
}

func newStr(a int, b int) *Str {
	return &Str{a: a, b: b}
}

func main() {

	type Str2 struct {
		Str
		c, d float32
	}
	s := Str2{Str{3, 4}, 1.0, 2.0}
	fmt.Println(s, s.Str, s.a)

	type per1 struct {
		name string
		sex  string
	}

	type per2 struct {
		name string
		age  int
	}

	type per struct {
		per1
		per2
	}

	str33 := newStr(10, 20)

	fmt.Println(str33)

	p := per{per1{name: "syh", sex: "男"}, per2{name: "syh2", age: 22}}

	//fmt.Println(p.name)   //错误的，编辑器能够识别这个错误--不能识别name是谁的name
	fmt.Println(p.sex, p.per2.name) //链式的获取，正确的
}
