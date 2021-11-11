package Figures

type Circle struct {
	center *Point
	radius float64
}

func (o *Circle) SetCenter(center *Point) {
	o.center = center
}

func (o *Circle) SetRadius(radius float64) {
	o.radius = radius
}

func (o *Circle) GetCenter() *Point {
	return o.center
}

func (o *Circle) GetRadius() float64 {
	return o.radius
}

func CreateCircle(center *Point, radius float64) *Circle {
	return &(Circle{center, radius})
}
