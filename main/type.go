package main

//
//import "fmt"
//
//// 将NewInt定义为int类型 这个是声明一个新的type
//type NewInt int
//
//// 将int取一个别名叫IntAlias
//type IntAlias = int
//
//func main() {
//
//	//type TypeAlias = Type   类型别名，给Type取一个别名
//	//虽然是一个别名，但TypeAlias本质上就是Type。
//
//	// 将a声明为NewInt类型
//	var a NewInt
//	// 查看a的类型名
//	fmt.Printf("a type: %T\n", a)
//	// 将a2声明为IntAlias类型
//	var a2 IntAlias
//	// 查看a2的类型名
//	fmt.Printf("a2 type: %T\n", a2)
//}
import (
	"fmt"
	"reflect"
)

// 定义商标结构
type Brand struct {
}

// 为商标结构添加Show()方法
func (t Brand) Show() {

}

// 为Brand定义一个别名FakeBrand
type FakeBrand = Brand

// 定义车辆结构
type Vehicle struct {
	// 嵌入两个结构
	FakeBrand
	Brand
}

func main() {
	// 声明变量a为车辆类型
	var a1 Vehicle

	// 指定调用FakeBrand的Show
	a1.FakeBrand.Show()
	// 取a的类型反射对象
	ta := reflect.TypeOf(a1)
	// 遍历a的所有成员
	for i := 0; i < ta.NumField(); i++ {
		// a的成员信息
		f := ta.Field(i)
		// 打印成员的字段名和类型
		fmt.Printf("FieldName: %v, FieldType: %v\n", f.Name, f.Type.Name())

	}
	type newType int
	type aliasType = int

	var a newType
	var b aliasType
	fmt.Printf("newType typeof: %T ,aliasType typeof: %T", a, b)
}
