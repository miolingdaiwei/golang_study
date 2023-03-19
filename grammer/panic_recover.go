package main

import (
	"fmt"
	"regexp"
)

func regPanic(str string) *regexp.Regexp {
	reg, err := regexp.Compile(str)
	if err != nil {
		panic(`regexp:Compile(` + regexp.QuoteMeta(str) + `):` + err.Error())
	}
	return reg
}

func main() {
	//i := 0
	//for {
	//	i++
	//	if i > 10 {
	//		panic("do panic")
	//	}
	//}
	//regPanic("str")

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("this is a word with defer&recover%s\n", err)
		}
	}()
	fmt.Println("this is a work without defer")

	panic("err")
}
