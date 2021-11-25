package main

import "fmt"

func main() {

	i := 2
	x := []int{3, 5, 7}
	for i, x[i] = range x {
		fmt.Printf("i = %d, x[i] = %d\n", i, x[i])
		break
	}

	fmt.Printf("i = %d, x = %v\n", i, x)
}
