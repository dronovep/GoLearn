package main

import "fmt"

var marr = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var sl = marr[2:7]
var sl2 = sl[0:6:7] // новый слайс может быть шире исходного, если позволяет capacity
//var sl3 = sl2[0:8]	// нельзя сделать слайс с length или capacity больше, чем capacity исходного
//var sl3 = sl2[0:5:8]	// нельзя сделать слайс с length или capacity больше, чем capacity исходного
var sl3 = sl2[0:5]

/*
Касательно length и capacity:
Length - показывает исключительно длину самого слайса. Слайс можно представить просто как рамку, которая
ездит по реальному массиву и может сжимать до 0 или растягиваться до размера массива
Capacity - длина от начала слайса до конца подлежащего массива

Как по факту работают Lenght и Capacity:
Length - ограничивает индексы, по которым можно обратиться к элементу слайса
Capacity - ограничивает Length и Capacity нового слайса, созданного на основе старого
*/

func main() {
	printSlice(sl)
	printSlice(sl2)
	printSlice(sl3)

	slength := 3
	sl4 := marr[:slength] // величина слайса может формироваться динамически
	printSlice(sl4)
	marr[0] = 256
	printSlice(sl4)

	//marr2 := [slength]int{}	// тогда как массив может быть объявлен с константным указанием размера
}

func printSlice(slice []int) {
	fmt.Printf("sl: len = %d, cap = %d, value = %v\n", len(slice), cap(slice), slice)
}
