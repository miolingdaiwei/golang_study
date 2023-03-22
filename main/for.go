package main

import "fmt"

func main() {

	//normal loop
	for i := 0; i < 100; i++ {
		fmt.Printf("i%d   ", i)
	}
	fmt.Println()

	//infinite loop  /   endless loop
	n := 0
	for {
		n++
		fmt.Println("this is do...")
		if n > 100 {
			fmt.Println("break...")
			break
		}
	}

	//
	var i int
	for {
		fmt.Printf("%d  ", i)
		if i > 10 {
			goto breakHere
		}
		i++
	}

breakHere:
	fmt.Println("break")

}
