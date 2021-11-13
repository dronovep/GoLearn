package main

///*
//#cgo LDFLAGS: -L/usr/lib64/libSDL-1.2.so.0 -l libSDL
//#include <SDL/SDL.h>
//*/
//import "C"
import (
	"fmt"
)

import "octadev.ru/GoLearn/Swpack"

func main() {

	//fmt.Printf("Cpack.Sum(100, 50) = %d\n", Cpack.Sum(12, 36))
	fmt.Printf("It works!\n")
	fmt.Printf("Swpack.wsum(15, 55) = %d\n", Swpack.Wsum(15, 55))

}
