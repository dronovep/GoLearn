package main

import "fmt"

var marr [5]int = [5]int{3, 2, 1, 6, 7} //изначальный массив
var sl []int = marr[2:4]                // слайс со 2го по 4й элемент (4й не включительно)

var sl2 []int = marr[1:5]

var sl3 []int = make([]int, 3, 6)

var sl4 []int = sl3[2:]

func main() {

	//sl3 = []int{3, 2, 1, 6, 7, 8}

	for i := 0; i < len(sl); i++ {
		fmt.Printf("sl[%d] = %d\n", i, sl[i])
	}

	for i := 0; i < len(marr); i++ {
		fmt.Printf("marr[%d] = %d\n", i, marr[i])
	}

	for i := 0; i < len(sl2); i++ {
		fmt.Printf("sl2[%d] = %d\n", i, sl2[i])
	}

	for i := 0; i < len(sl3); i++ {
		fmt.Printf("sl3[%d] = %d\n", i, sl3[i])
	}

	for i := 0; i < len(sl4); i++ {
		fmt.Printf("sl4[%d] = %d\n", i, sl4[i])
	}
}
