package main

import (
	"fmt"
	"sync"
)

func main() {
	//var map2 map[string]string  //声明,并没有生成实例
	map1 := make(map[string]string) //make构造，返回map实例
	map1["key1"] = "value1"
	map1["key2"] = "value2"
	map1["key3"] = "value3"
	fmt.Println(map1)
	for s, s2 := range map1 {
		fmt.Println(s, s2)
	}
	//map[key1:value1 key2:value2 key3:value3]
	//key1 value1
	//key2 value2
	//key3 value3

	var scene sync.Map
	scene.Store("1", 1)
	scene.Store("2", "2")
	scene.Store(3, "v3")
	scene.Store(4, 4)
	//println(syncmap1.Load("key1"))
	fmt.Println(scene.Load("2"))
	fmt.Println(scene.LoadAndDelete(4))
}
