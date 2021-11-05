package main

import "fmt"

/**
Для получения мнимой части комплексного числа просто добавляем "i" к концу вещественного числа
*/

var (
	img1 complex64 = 35.5 + 25.3e-4i
)

func main() {
	mvar := 35
	fmt.Printf(
		"%f, %d",
		img1,
		mvar,
	)
}
