package figures

import "math"

type Point struct {
	x float64
	y float64
}

func (o *Point) GetX() float64 {
	return o.x
}

func (o *Point) GetY() float64 {
	return o.y
}

func (o *Point) SetCoordinates(x float64, y float64) {
	o.x = x
	o.y = y
}

func (o *Point) GetLength() float64 {
	return math.Sqrt(o.x*o.x + o.y*o.y)
}

func CreatePoint(x float64, y float64) *Point {
	return &Point{x, y}
}
