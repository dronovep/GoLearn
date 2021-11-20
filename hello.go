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

func (o *Point) getLength() float64 {
	return math.Sqrt(o.x*o.x + o.y*o.y)
}

func (o *Point) getX() float64 {
	return o.x
}

func (o *Point) getY() float64 {
	return o.y
}

func (o *Point) setCoordinates(x float64, y float64) {
	o.x = x
	o.y = y
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

// Пример с включением одних интерфейсов в другой
type ICanGreet interface {
	greet() string
}

type ICanBye interface {
	bye() string
}

type ICanMeet interface {
	ICanGreet
	ICanBye
}

type Dog struct {
	name string
}

func (o *Dog) greet() string {
	return fmt.Sprintf("Woof!Woof! My name is %s", o.name)
}

func (o *Dog) bye() string {
	return fmt.Sprintf("Woof!Bye!")
}

func main() {
	p := &(Point{3, 2})
	var d IHaveCoordinates = &(Point{17, 33})
	testLength(p)
	testCoordinates(p)
	// а вот testLength(d) тут уже не вызвать, так как d был явно указан тип интерфейса, а не Point
	testCoordinates(d)

	var meeter ICanMeet = &Dog{"Fido"}
	fmt.Printf("\n\nПриветствие: %s\n", meeter.greet())
	fmt.Printf("Прощание: %s\n", meeter.bye())
}

/**
Как можно заметить, методы объявляются вне структуры, а не внутри нее.
На принадлежность метода к структуре указывает ресивер метода, который должен быть типом этой структуры
или указателем на эту структуру
*/

/**
Важно!!! Поскольку набор методов для структуры может быть двояк (методы с ресиверами-объектами и ресиверами-указателями)
Нужно четко понимать, какой интерфейс реализуется структурой и ее методами:
Указатель на структуру будет обладать методами с обоими типами ресиверов
объект структуры будет обладать только методами с ресивером-объектом
Поэтому возможна такая ситуация, которая приведет к ошибке компиляции:
Мы попытаемся присвоить переменной типа интерфейса, объект структуры, но при этом у структуры недостаточно методов
с ресивером-объектом, для реализации интерфейса (даже если абсолютно все есть с ресивером-указателем)
Тогда присвоение будет невозможно. Для примера смотри ветку "Клонирование"
*/
