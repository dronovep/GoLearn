package main

import (
	"fmt"
	"octadev.ru/GoLearn/Compound"
	"octadev.ru/GoLearn/Interfaces"
	"octadev.ru/GoLearn/figures"
)

// Важно!!! Обрати внимание, что здесь не нужно указывать указатель на фигуру
// несмотря на то, что ты по факту передаешь указатель на объект фигуры - он все равно реализует нужный интерфейс
// так как методы работают с указателями
func printFigure(o Interfaces.Figure) {
	fmt.Printf("Printing Figure\n")
	fmt.Printf("Figure: x = %f, y = %f, length = %f\n", o.GetX(), o.GetY(), o.GetLength())
}

func printTheFigure(o Interfaces.IamFigure) {
	fmt.Printf("Printing TheFigure\n")
	fmt.Printf("Figure: x = %f, y = %f, length = %f\n", o.GetX(), o.GetY(), o.GetLength())
}

func main() {
	circle := figures.CreateCircle(12, 17, 24)
	fmt.Printf("circle: x = %f, y = %f, radius = %f\n", circle.GetX(), circle.GetY(), circle.GetRadius())

	circle.SetCoordinates(1.5, 3.4)
	circle.SetRadius(25.0)

	fmt.Printf("circle: x = %f, y = %f, radius = %f\n", circle.GetX(), circle.GetY(), circle.GetRadius())
	// GetLength был переопределен для Circle - здесь выдастся длина окружности
	fmt.Printf("Circle length = %f\n", circle.GetLength())
	// Для вызова GetLength чисто поинта можно использовать такой синтаксис
	fmt.Printf("Circle distance = %f\n", circle.Point.GetLength())
	fmt.Printf("Circle distance = %f\n", circle.GetDistance())

	fmt.Printf("теперь работаем с Compound\n")
	cmp := Compound.CreateCompound()
	//fmt.Printf("Работаем с валютой %s\n", cmp.GetValute()) так вызвать не получится, Go не знает какой из 2х методов вызывать
	fmt.Printf("Работаем с рублевой валютой %s\n", cmp.Ruble.GetValute())
	fmt.Printf("Работаем с долларовой валютой %s\n", cmp.Dollar.GetValute())

	//Не получится вызвать неоднозначный метод, даже если определить собственный метод
	//fmt.Printf("Работаем с композитной валютой %s\n", cmp.GetValute())
	// собственный метод даже вот так не вызвать
	//fmt.Printf("Работаем с композитной валютой %s\n", cmp.Compound.GetValute())

	fmt.Printf("\n\nТеперь поработаем с интерфейсами\n")
	printFigure(circle)
	printTheFigure(circle)
}
