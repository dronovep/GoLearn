package Cpack

/*
#include "mcsum.h"
*/
import "C"

func Sum(a int, b int) int {
	return int(C.sum(C.int(a), C.int(b)))
}
