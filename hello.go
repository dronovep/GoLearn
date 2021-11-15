package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (o *Point) getX() float64 {
	return o.x
}

func (o *Point) getY() float64 {
	return o.y
}

func (o *Point) getLength() float64 {
	return math.Sqrt(o.x*o.x + o.y*o.y)
}

func (o *Point) setCoords(x float64, y float64) {
	o.x = x
	o.y = y
}

// Возвращенный из функции объект -  к сожалению от него автоматически взять адрес для метода с указателем
//func CreatePoint(x float64, y float64) Point {
//	return Point{x, y}
//}

//Зато если мы сами вручную возвращаем адрес из метода - все работает
func CreatePoint(x float64, y float64) *Point {
	return &Point{x, y}
}

func main() {
	// Возвращенный из функции объект -  к сожалению от него автоматически взять адрес для метода с указателем
	//fmt.Printf("Dynamic Point: length = %f\n", CreatePoint(3, 2).getLength())

	// И даже вот так не получится. ругнется, что нельзя получить адрес, возвращенного из метода объекта
	//mgetLength := CreatePoint(3, 2).getLength
	//fmt.Printf("Dynamic Point: length = %f\n", mgetLength())

	//Зато если мы сами вручную возвращаем адрес из метода - все работает
	fmt.Printf("Dynamic Point: length = %f\n", CreatePoint(3, 2).getLength())
	fmt.Printf("Dynamic Point: length = %f\n", CreatePoint(15, 25).getLength())
}
