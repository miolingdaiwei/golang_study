package main

import "fmt"

func main() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	//3
	//2
	//1
}
