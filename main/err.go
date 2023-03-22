/**
 * @Author: lingjiang
 * @Description:
 * @File:  err
 * @Version: 1.0.0
 * @Date: 2023/03/21 11:13
 */

package main

import (
	"fmt"
	"math"
)

type dualError struct {
	Num     float64
	problem string
}

func (e dualError) Error() string {
	return fmt.Sprintf("Wrong!!!,because \"%f\" is a negative number", e.Num)
}
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return -1, dualError{Num: f}
	}
	return math.Sqrt(f), nil
}
func main() {
	result, err := Sqrt(-13)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
