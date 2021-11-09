package main

import "fmt"
import "math"

type Point struct {
	x        int
	y        int
	distance func() float64
}

var p Point = Point{3, 2, func() float64 {
	return math.Sqrt(x*x + y*y)
}}

func main() {
	fmt.Printf("Привет мир!\n")
	fmt.Printf("Точка: x = %d, y = %d\n", p.x, p.y)
}
