package main

import (
	"fmt"
	"testing"
)

type Point struct {
	x            float64
	y            float64
	setCoordsAlg func(o *Point, x float64, y float64)
}

func (o *Point) setCoords(x float64, y float64) {
	o.setCoordsAlg(o, x, y)
}

func StandardSetCoords(o *Point, x float64, y float64) {
	o.x = x
	o.y = y
}

func setCoordsWithTen(o *Point, x float64, y float64) {
	o.x = x + 10
	o.y = y + 10
}

func main() {
	p := Point{0, 0, StandardSetCoords}

	p.setCoords(12, 24)

	fmt.Printf("p = %v\n", p)
	p.setCoordsAlg(&p, 35, 84)
	fmt.Printf("p = %v\n", p)
}

func BenchmarkStandardSetCoords(t *testing.B) {
	p := Point{0, 0, standardSetCoords}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for i := 0; i < 100; i++ {
			p.setCoordsAlg(&p, float64(i), 0)
		}
	}
}
