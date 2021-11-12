package main

import "fmt"

// вне скобкового блока const значение iota не инкрементируется
const alpha = iota
const beta = iota
const gamma = iota
const delta = iota

//внутри скобкового блока const значение iota инкрементируется
const (
	mother  = iota
	father  = iota
	sister  = iota
	brother = iota
)

// можно даже не всем константам указывать iota, достаточно на самом деле только первой, остальные получат инкремент
const (
	Kia = iota
	Chevrolet
	Audi = iota
	BMW
	Toyota   = 15 //здесь интересный момент, тойота получила не йоту, но само значение йота инкрементировалось
	Cadillac = iota
)

// В скобковом блоке const , iota начинает инкрементироваться сразу по каждой объявляемой константе и неважно какого типа эта константа (в том числе и не типа iota)
const (
	Morrowind uint    = 15
	Angle     float64 = 3.3

	Vasya  = "Voronov"
	Number = iota
)

//значение iota инкрементируется только при вводе новой константы, но не от самого использования iota
const (
	A = iota + iota + iota
	B = iota
)

/**
Получается в языке Go нет как такового Enum-а, вместо этого можно вынести набор констант в отдельный package,
дать им значения iota и экспортировать, возможно даже написав к ним парочку функций
*/

func main() {
	// вне скобкового блока const значение iota не инкрементируется
	const mconst = iota
	fmt.Printf("alpha = %d\n", alpha)
	fmt.Printf("beta = %d\n", beta)
	fmt.Printf("gamma = %d\n", gamma)
	fmt.Printf("delta = %d\n", delta)
	fmt.Printf("mconst = %d\n", mconst)

	fmt.Printf("\n\n")

	//внутри скобкового блока const значение iota инкрементируется
	fmt.Printf("mother = %d\n", mother)
	fmt.Printf("father = %d\n", father)
	fmt.Printf("sister = %d\n", sister)
	fmt.Printf("brother = %d\n", brother)

	fmt.Printf("\n\n")

	// можно даже не всем константам указывать iota, достаточно на самом деле только первой, остальнгые получат инкремент
	fmt.Printf("Kia = %d\n", Kia)
	fmt.Printf("Chevrolet = %d\n", Chevrolet)
	fmt.Printf("Audi = %d\n", Audi)
	fmt.Printf("BMW = %d\n", BMW)
	fmt.Printf("Toyota = %d\n", Toyota) //здесь интересный момент, тойота получила не йоту, но само значение йота инкрементировалось
	fmt.Printf("Cadillac = %d\n", Cadillac)

	fmt.Printf("\n\n")

	// В скобковом блоке const , iota начинает инкрементироваться сразу по каждой объявляемой константе и неважно какого типа эта константа (в том числе и не типа iota)
	fmt.Printf("Morrowind = %d\n", Morrowind)
	fmt.Printf("Angle = %f\n", Angle)
	fmt.Printf("Vasya = %s\n", Vasya)
	fmt.Printf("Number = %d\n", Number)

	fmt.Printf("\n\n")

	//значение iota инкрементируется только при вводе новой константы, но не от самого использования iota
	fmt.Printf("A = %d\n", A)
	fmt.Printf("B = %d\n", B)
}
