package main

import "fmt"

/** внутри `` - сырые строки
внутри "" - интерпретируемые строки (в них обрабатываются спецсимволы)
*/

var (
	family    = `mother father sister brother\nYes!`
	family_i  = "mother father sister brother\nYes!"
	family_r  = `мама папа сестра брат \nДа!`
	family_ri = "мама папа сестра брат \nДа!"
	mama      = "mam\u0041\n"
)

func main() {
	fmt.Printf("Family is %s, Again Family is %s", family, family_i)
	fmt.Printf("\n\n\n")
	fmt.Printf("Семья это %s, Опять семья это %s", family_r, family_ri)
	fmt.Printf(`Вот это \r`) // проверяем поведение \r в сырой строке
	fmt.Printf("%s", mama)
}
