package main

import "fmt"
import "octadev.ru/GoLearn/msupport"

type mtype struct {
	maka int
	vala int
}

var mstr mtype = mtype{5, 8}
var p struct {
	x int
	y int
} = struct {
	x int
	y int
}{1, 7}

type (
	mint  int
	mfunc func(a int, b int) string
	rstr  func() string
)

var mfunction mfunc = func(a int, b int) string {
	return fmt.Sprintf("mfunction: a = %d, b = %d", a, b)
}

func foo(method rstr) {
	fmt.Printf("%s\n", method())
}

func vaka() string {
	return "Returning vaka"
}

func main() {

	//for i := 0; i < len(slice); i++ {
	//	fmt.Printf("sl[%d] = %d\n", i, slice[i])
	//}

	//fmt.Printf("%s\n", mstr)
	fmt.Printf("first_component = %d, second_component = %d\n", mstr.maka, mstr.vala)
	fmt.Printf("first_component = %d, second_component = %d\n", p.x, p.y)

	mstr := mfunction(msupport.Alpha, 101)
	fmt.Printf("%s\n", mstr)
	foo(vaka)
}
