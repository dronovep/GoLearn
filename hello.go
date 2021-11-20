package main

import (
	"fmt"
	"math"
)

type Clonable interface {
	clone() Clonable
}

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

//метод clone приходится делать с ресивером-объектом, иначе была бы ошибка реализации интерфейса
// (ну или методу с ресивером указателем пришлось бы возвращать указатель, тогда само клонирование уже не происходило бы)
func (o Point) clone() Clonable {
	return o
}

func main() {
	orig := Point{3, 2}
	dub := orig.clone().(Point)
	dub.setCoords(15, 25)

	fmt.Printf("orig Point: x= %f, y = %f\n", orig.getX(), orig.getY())
	fmt.Printf("dub Point: x= %f, y = %f\n", dub.getX(), dub.getY())
}
