package main

type Point struct {
	x            float64
	y            float64
	SetCoordsAlg func(o *Point, x float64, y float64)
}

func (o *Point) SetCoords(x float64, y float64) {
	o.SetCoordsAlg(o, x, y)
}

func StandardSetCoords(o *Point, x float64, y float64) {
	o.x = x
	o.y = y
}

func SetCoordsWithTen(o *Point, x float64, y float64) {
	o.x = x + 10
	o.y = y + 10
}
