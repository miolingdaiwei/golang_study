package main

import "fmt"

func namedRetValues() (a, b int) {
	a = 1
	b = 2
	return
}

func doubleRe(a int, b string) (i int, j string) {
	return a, b
}
func main() {
	fmt.Println(namedRetValues())
	fmt.Println(doubleRe(2, "3"))

	f := func(word string) {
		fmt.Println("word is : ", word)
	}

	f("niganma")
}
