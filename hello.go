package main

import "fmt"

func msum(a int, b int) int {
	return a + b
}

func mmul(a int, b int) int {

	return a * b
}

var (
	mother  string = "Irina"
	father  string = "Pavel"
	brother string = "Anton"
	wife    string = "Dasha"
	son     string = "Cyrill"
	sson    string = "Alexey"
)

//var sym rune = '\u0062'
var code int = 333_157

func main() {
	//var m, f, s, b int = 5, 6, 7, 0xf

	//fmt.Printf("Hello dude! m + f = %d, s * b = %d\n", msum(m ,f), mmul(s, b))
	fmt.Printf("Code = %d", code)
	//println(m, f, s, b)
	//fmt.Printf(
	//	"Hello %s, %s, %s, %s, %s, %s",
	//		mother,
	//		father,
	//		brother,
	//		wife,
	//		son,
	//		sson,
	//	)
	//var a int = 3
	//var b int = 5
	//var p *int = &a
	//
	//fmt.Printf("p points to a, *p = %d, ", *p)
	//p = &b
	//fmt.Printf("p points to b, *p = %d", *p)
	//fmt.Printf("Blank identifier equals to %d", _)
}
