package main

import (
	"fmt"
	"math"
)

type IHaveLength interface {
	getLength() float64
}

type IHaveCoordinates interface {
	getX() float64
	getY() float64
	setCoordinates(x float64, y float64)
}

type Point struct {
	x float64
	y float64
}

func (this *Point) getLength() float64 {
	return math.Sqrt(this.x*this.x + this.y*this.y)
}

func (this *Point) getX() float64 {
	return this.x
}

func (this *Point) getY() float64 {
	return this.y
}

func (this *Point) setCoordinates(x float64, y float64) {
	this.x = x
	this.y = y
}

func testLength(obj IHaveLength) {
	fmt.Printf("obj length = %f\n", obj.getLength())
}

func testCoordinates(obj IHaveCoordinates) {
	fmt.Printf("obj X coordinate = %f\n", obj.getX())
	fmt.Printf("obj Y coordinate = %f\n", obj.getY())
	fmt.Printf("setting new coordinates\n")
	obj.setCoordinates(15, 25)
	fmt.Printf("obj coordinates now are: x = %f, y = %f\n", obj.getX(), obj.getY())
}

func main() {
	p := &(Point{3, 2})
	testLength(p)
	testCoordinates(p)
}

/**
Важный момент! Как можно заметить, методы объявляются вне структуры, а не внутри нее.
На принадлежность метода к структуре указывает ресивер метода, который должен быть типом этой структуры
или указателем на эту структуру
*/
