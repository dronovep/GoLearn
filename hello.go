package main

import "fmt"

var flag bool = true
var chislo uint8 = 128

func main() {
	fmt.Printf("flag = %v, !flag = %v\n", flag, !flag)
	fmt.Printf("chislo = %v, !chislo = %v\n", chislo, ^chislo)
}
