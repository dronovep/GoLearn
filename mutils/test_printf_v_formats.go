package mutils

import (
	"fmt"
	. "octadev.ru/GoLearn/mtypes"
)

func TestPrintfVformats() {
	me := Person{Name: "Евгений", Surname: "Дронов"}
	// {Евгений Дронов}
	fmt.Printf("%v\n", me)

	//  "+" - позволяет выводить не только значения полей, но и имена   {name:Евгений surname:Дронов}
	fmt.Printf("%+v\n", me)

	// main.Person{name:"Евгений", surname:"Дронов"}  здесь еще выводится пакет и тип
	fmt.Printf("%#v\n", me)

	// main.Person  - выводится только пакет и тип, без значения
	fmt.Printf("%T\n", me)
}
