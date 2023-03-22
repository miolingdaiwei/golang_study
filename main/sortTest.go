package main

import (
	"fmt"
	"sort"
)

// MyStringList 将[]string定义为MyStringList类型
type MyStringList []string

// Len 实现sort.Interface接口的获取元素数量方法
func (m MyStringList) Len() int {
	return len(m)
}

// Less 实现sort.Interface接口的比较元素方法
func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

// Swap 实现sort.Interface接口的交换元素方法
func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

type HeroKind int

type Hero struct {
	Name string
	Kind HeroKind
}

func main() {
	// 准备一个内容被打乱顺序的字符串切片
	names := MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}

	// 使用sort包进行排序
	sort.Sort(names)
	// 遍历打印结果
	for _, v := range names {
		fmt.Printf("%s\n", v)
	}

	const (
		fir = iota
		sec
		thi
		fou
		fiv
	)

	arr := []*Hero{
		{"men", fir},
		{"men2", fou},
		{"men4", sec},
		{"men3", fiv},
		{"men5", thi},
	}
	sort.Slice(arr, func(i, j int) bool {
		// 编写实现 less 接口
		if arr[i].Kind != arr[j].Kind {
			return arr[i].Kind < arr[j].Kind
		}
		return arr[i].Name < arr[j].Name
	})

	for _, v := range arr {
		fmt.Println(*v)
	}
}
