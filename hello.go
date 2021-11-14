package main

import "fmt"

var mmap = map[string]int{
	"mother":  1,
	"father":  2,
	"sister":  3,
	"brother": 4,
}

type Car struct {
	brand string
	model string
}

//Даже если это не мап, имена полей все равно можно использовать как ключи
//Но один из двух вариантов 1) или всем полям даются ключи, либо ни одному - но тогда идут в том порядке, как в структуре
var mycar = Car{
	brand: "Kia",
	model: "Sorento",
}

var marray = make([]int, 5)

// Здесь объявляется массив, а не слайс. многоточие испольуется, чтобы не высчитывать самому длинну массива
// а компилятор посчитал ее сам на основе переданных инициализатору значений
var marr = [...]string{"mother", "father", "sister", "brother"}

type Point struct {
	x float64
	y float64
}

//Важно!!! Обе записи имеют право на жизнь
var mtest = [...]Point{Point{12, 24}, Point{1, 3}}
var mtest2 = [...]Point{{12, 24}, {1, 3}}

func main() {
	for key := range mmap {
		fmt.Printf("%s has value %d\n", key, mmap[key])
	}

	fmt.Printf("\nМоя машина - %s %s\n", mycar.brand, mycar.model)

	for i := range marr {
		fmt.Printf("marr[%d] = %s\n", i, marr[i])
	}

	for i := range mtest {
		fmt.Printf("mtest[%d] = Point(%f, %f)\n", i, mtest[i].x, mtest[i].y)
	}

	for i := range mtest2 {
		fmt.Printf("mtest2[%d] = Point(%f, %f)\n", i, mtest2[i].x, mtest2[i].y)
	}
}
