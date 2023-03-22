package main

import (
	"log"
	"runtime"
	"time"
)

type Road int

func findRoad(r *Road) {
	log.Println("road:", *r)
}

func entry() {
	var rd = Road(999)
	r := &rd
	runtime.SetFinalizer(r, findRoad)
	//	设置这个函数被回收
}

func main() {
	entry()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		//time.sleep   time.second 1000
		runtime.GC()
		//  GC
	}
}
