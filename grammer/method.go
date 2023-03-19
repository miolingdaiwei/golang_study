package main

import "fmt"

// Person 定义接收者结构体
type Person struct {
	name string
	age  int
}

// func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
//函数体
//}
func (p Person) pMSG() { //这里的意义就是Person这个结构体接收到了pMSG这个方法
	fmt.Printf("this person's name is: %s and age is %d\n", p.name, p.age)
}

//定义多个方法到接收者上
func (p Person) pName() {
	fmt.Println("name is " + p.name)
}

func main() {
	p := Person{
		name: "mio",
		age:  18,
	}
	//p.就可以看到Person结束到的方法
	p.pMSG()
	p.pName()
}
