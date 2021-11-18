package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (o *Point) pGetX() float64 {
	fmt.Printf("Invoking Point.pGetX()\n")
	return o.x
}

func (o *Point) pGetY() float64 {
	fmt.Printf("Invoking Point.pGetY()\n")
	return o.y
}

func (o Point) oGetX() float64 {
	fmt.Printf("Invoking Point.oGetX()\n")
	return o.x
}

func (o Point) oGetY() float64 {
	fmt.Printf("Invoking Point.oGetY()\n")
	return o.y
}

func (o *Point) pSetCoords(x float64, y float64) {
	fmt.Printf("Invoking Point.pSetCoords(%f, %f)\n", x, y)
	o.x = x
	o.y = y
}

func (o Point) oSetCoords(x float64, y float64) {
	fmt.Printf("Invoking Point.oSetCoords(%f, %f)\n", x, y)
	o.x = x
	o.y = y
}

func (o *Point) pGetLength() float64 {
	fmt.Printf("Invoking Point.pGetLength()\n")
	return math.Sqrt(o.x*o.x + o.y*o.y)
}

func (o Point) oGetLength() float64 {
	fmt.Printf("Invoking Point.oGetLength()\n")
	return math.Sqrt(o.x*o.x + o.y*o.y)
}

// Теперь поработаем с интерфейсом
type IHaveCoords interface {
	getX() float64
	getY() float64
}

type ICanSetCoords interface {
	setCoords(x float64, y float64)
	oSetCoords(x float64, y float64)
	vaka(s string)
}

type IHaveLength interface {
	getLength() float64
}

type iPoint struct {
	x float64
	y float64
}

func (o *iPoint) getX() float64 {
	return o.x
}

func (o *iPoint) getY() float64 {
	return o.y
}

func (o *iPoint) setCoords(x float64, y float64) {
	o.x = x
	o.y = y
}

func (o iPoint) oSetCoords(x float64, y float64) {
	o.x = x
	o.y = y
}

func (o iPoint) vaka(s string) {
	fmt.Printf("Invoking vaka with s = %s\n", s)
	//o.x = x

}

func (o *iPoint) getLength() float64 {
	return math.Sqrt(o.x*o.x + o.y*o.y)
}

func main() {
	mpoint := Point{4, 3}
	fmt.Printf("mpoint: x = %f, y = %f, length = %f\n", mpoint.pGetX(), mpoint.pGetY(), mpoint.pGetLength())

	// Это не сработает, причем именно в runtime
	// Мы выбрали метод pSetCoords, у которого ресивер  - указатель, а в выражении метода перед селектором стоит объект, а не указатель
	//pset := Point.pSetCoords // не сработает именно эта строчка (invalid method expression Point.pSetCoords (needs pointer receiver: (*Point).pSetCoords))
	//pset(mpoint, 12, 24)
	//fmt.Printf("mpoint: x = %f, y = %f, length = %f\n", mpoint.pGetX(), mpoint.pGetY(), mpoint.pGetLength())

	// а вот так сработает
	pset := (*Point).pSetCoords // мы явно должны в селекторе указать тип ресивера, совместимый с выбранным методом
	pset(&mpoint, 12, 24)
	fmt.Printf("mpoint: x = %f, y = %f, length = %f\n", mpoint.pGetX(), mpoint.pGetY(), mpoint.pGetLength())

	// вариант с ресивером - объектом
	// в данном случае мы не увидим изменений в объекте mpoint, так как обрабатывается по факту его клон
	oset := Point.oSetCoords
	oset(mpoint, 1, 2)
	fmt.Printf("mpoint: x = %f, y = %f, length = %f\n", mpoint.pGetX(), mpoint.pGetY(), mpoint.pGetLength())

	// Хотя задать ресивер указатель для метода с ресивером - объектом можно
	// Получается сделать ресивер-указатель для метода-объекта можно, а вот наоборот
	// Сделать ресивер-объект для метода - указателя нельзя
	// Важно!!! можно то можно, вот только по факту работать это будет все равно так, словно вызывался метод с ресивером-объектом
	// Проверь вывод кода ниже и поймешь
	fmt.Printf("Last experiment\n\n")
	nset := (*Point).oSetCoords
	nset(&mpoint, 33, 58)
	fmt.Printf("mpoint: x = %f, y = %f, length = %f\n", mpoint.pGetX(), mpoint.pGetY(), mpoint.pGetLength())

	//Теперь поработаем с интерфейсом
	mpi := iPoint{22, 34}
	//mscoords := ICanSetCoords.setCoords
	//mscoords(&mpi,3, 4)
	//fmt.Printf("after mscoords iPoint: x = %f, y = %f, length = %f\n", mpi.getX(), mpi.getY(), mpi.getLength())

	//ВАЖНО!!! крайне странный момент языка. В сигнатуре метода, объявляемого в интерфейсе, ведь ничего не говорится о типе ресивера
	//Складывается впечатление, что если мы используем объект через интерфейс, то язык ожидает в любом методе тот тип ресивера, который был
	// в первом попавшемся или вроде того
	oscoords := ICanSetCoords.vaka
	oscoords(mpi, "wrfrf")
	fmt.Printf("after mscoords iPoint: x = %f, y = %f, length = %f\n", mpi.getX(), mpi.getY(), mpi.getLength())

	//fmt.Printf("Now working with Interface\n\n")
	//setCoords := ICanSetCoords.oSetCoords
	//setCoords(mpoint, 15, 35)
}
