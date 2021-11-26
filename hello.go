package main

import (
	"fmt"
)

func main() {
	//var a interface{} = 32
	//var a interface{} = 7.3
	var a interface{} = true

	if _, isInt := a.(int); isInt {
		fmt.Printf("Вы задали целое число %d\n", a)
	} else if _, isFloat := a.(float64); isFloat {
		fmt.Printf("Вы задали вещественное число %f\n", a)
	} else {
		fmt.Printf("Вы задали значение %v\n", a)
	}

}
