package main

import "fmt"

var (
	a       int = 'a'
	A       int = 'A'
	newline     = '\n'
	eof         = '\u0000'
	uni_A       = '\u0041' // ASCII код символа A - это 65, в юникоде U + 65 в 16ричной системе счисления - U+0041
)

func main() {
	fmt.Printf(
		"%d\n"+
			"%d\n"+
			"%d\n"+
			"%d\n"+
			"%c\n",
		a,
		A,
		newline,
		eof,
		uni_A,
	)
}
