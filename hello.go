package main

import (
	"fmt"
	"reflect"
)

type (
	int1 = int
	int2 = int
	int3 = int2

	float1 = float64
	float2 = float64

	str1 = string
	str2 = string

	run1 = rune
	run2 = rune

	byte1 = byte
	byte2 = byte

	slt1 = []int
	slt2 = []int

	trp1 = [3]float64
	trp2 = [3]float64

	point1 = struct {
		x float64
		y float64
	}
	point2 = struct {
		x float64
		y float64
	}
	point3 = point2
	rpoint = struct {
		a float64
		b float64
	}
)

func printInfo(value interface{}) {
	fmt.Printf("value = %v, type of value is %s\n", value, reflect.TypeOf(value))
}

func main() {
	//var a int1 = 3
	//var b int2 = a
	//printInfo(b)

	//var a float1 = 3.0
	//var b float2 = a
	//printInfo(b)

	//var a int1 = 15
	//var b int3 = a
	//printInfo(b)

	//var a str1 = "mother"
	//var b str2 = a
	//printInfo(b)

	//var a run1 = 'e'
	//var b run2 = a
	//printInfo(b)

	//var a byte1 = 60
	//var b byte2 = a
	//printInfo(b)

	//var a = [5]int{0, 1, 2, 3, 4}
	//var b slt1 = a[:3]
	//var c slt2 = b
	//printInfo(c)

	//var a trp1 = trp1{3.0, 2.0, 1.15}
	//var b trp2 = a
	//printInfo(b)

	//var a point1 = point1{3, 2}
	//var b point2 = a
	//var c point3 = b
	//printInfo(b)
	//printInfo(c)

	/*
		Все закомvенченые выше операторы работают
		Значит отдельно определенные типы одинаковы и взаимозаменяемы если каждый состоит из одного и того же БАЗОВОГО
		типа
	*/

	////var a = struct{x float64; y float64}{15, 5} // так нельзя, у компонентов точки другие имена - уже другой тип
	//var a = struct{a float64; b float64}{15, 5} //так можно
	////var b rpoint = rpoint(a)	// и даже явное приведение типа не поможет - нельзя сконвертировать
	//var b rpoint = a
	//printInfo(b)

	//var mstr string = "father"
	//var pmstr string = mstr[2:]	// слайс от строки - тоже строка
	//printInfo(pmstr)
}
