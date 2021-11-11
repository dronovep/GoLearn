package main

import (
	"fmt"
	"octadev.ru/GoLearn/Figures"
)

func main() {
	p := Figures.CreatePoint(12, 7)
	fmt.Printf("p Point: x = %f, y = %f, length = %f\n", p.GetX(), p.GetY(), p.GetLength())
	p.SetCoordinates(3, 2)
	fmt.Printf("p Point: x = %f, y = %f, length = %f\n", p.GetX(), p.GetY(), p.GetLength())
	d := Figures.CreatePoint(38, 37)
	fmt.Printf("d Point: x = %f, y = %f, length = %f\n", d.GetX(), d.GetY(), d.GetLength())
	fmt.Printf("p Point: x = %f, y = %f, length = %f\n", p.GetX(), p.GetY(), p.GetLength())
	fmt.Printf("p pointer = %p\n", p)
	fmt.Printf("d pointer = %p\n", d)

	//теперь создадим круг

	//круг возможно создать прямо на месте через композитные литералы, даже еслиих нужно вкладывать
	//позаботьтесь только о том, чтобы тогда поля структур  были экспортированы
	//circle := Figures.Circle{
	//	&Figures.Point{
	//		3,
	//		2,
	//	},
	//	10,
	//}
	//
	//fmt.Printf("Circle: center - Point: x = %f, y = %f, radius = %f\n", circle.center.GetX(), circle.center.GetY(), circle.radius)

	circle1 := Figures.CreateCircle(Figures.CreatePoint(15, 25), 64)
	fmt.Printf("Circle1: center - Point: x = %f, y = %f, radius = %f\n", circle1.GetCenter().GetX(), circle1.GetCenter().GetY(), circle1.GetRadius())
	circle2 := Figures.CreateCircle(Figures.CreatePoint(4, 3), 12)
	fmt.Printf("Circle2: center - Point: x = %f, y = %f, radius = %f\n", circle2.GetCenter().GetX(), circle2.GetCenter().GetY(), circle2.GetRadius())
	fmt.Printf("Circle1: center - Point: x = %f, y = %f, radius = %f\n", circle1.GetCenter().GetX(), circle1.GetCenter().GetY(), circle1.GetRadius())
}
