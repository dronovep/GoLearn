package main

import "fmt"

func main() {
	var a interface{} = 35.5

	switch a.(type) {
	case int:
		fmt.Printf("Значение типа int = %d\n", a)
	case float64:
		fmt.Printf("Значение типа float64 = %f\n", a)
	case bool:
		fmt.Printf("Значение типа bool = %v\n", a)
	default:
		fmt.Printf("Значение какого-то типа = %v\n", a)
	}

	fmt.Printf("a = %v\n", a)

	switch b := 24; b {
	case 24, 48:
		fmt.Printf("b = %d\n", b)
		fallthrough
	case 35:
		fmt.Printf("b = 35\n")
		fallthrough
	default:
		fmt.Printf("b чему-то там равно\n")
	}

}
