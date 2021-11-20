package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var mar = [10]int{3, 2, 1, 6, 7, 4, 8, 5, 9, 0}

var mstr = "mother_father_sister_brother"

var mmap = map[string]int{
	"mother":  1,
	"father":  2,
	"sister":  3,
	"brother": 4,
	"me":      5,
}

var nmap = map[int]string{
	0: "mother",
	1: "father",
	2: "brother",
	3: "me",
}

func main() {
	//Обратиться к массиву по индексу с не целочисленным типом не получится, даже если само число целое
	//var x float64 = 3
	var x int = 3

	//слайс, начинающийся с индекса 3 исходного массива и до конца
	sl := mar[3:]

	fmt.Printf("mar[%d] = %d\n", x, mar[x])
	fmt.Printf("\nРаботаем со слайсом sl\n")
	for i := range sl {
		fmt.Printf("mar[%d] = %d\n", i, sl[i])
	}
	fmt.Printf("\nРаботаем с указателем на массив\n")

	pmar := &mar
	//for i := range pmar {		// можно как pmar, так и *pmar
	for i := range *pmar {
		fmt.Printf("pmar[%d] = %d\n", i, pmar[i])
	}

	fmt.Printf("\nПеречисляем строку\n")
	for letter := range mstr {
		fmt.Printf("letter = %c\n", mstr[letter])
	}

	fmt.Printf("\nДелаем слайс строки\n")
	// Отдельный символ строки имеет тип byte, но подстрока - это не слайс по байтам, а тоже строка
	var ssl string = mstr[:10]
	tp := reflect.TypeOf(ssl)
	fmt.Printf("ssl type is %s\n", tp.Name())
	for letter := range ssl {
		fmt.Printf("letter = %c\n", ssl[letter])
	}

	fmt.Printf("\nработаем с мапами\n")
	// Обратимся к несуществующему элементу мапы - должны получить "нулевое" значение подлежащего типа
	// Также при обращении к мапе всегда неявно отдается второе значение, показывающее был ли элемент по указанному ключу
	result, found := mmap["uncle"]
	fmt.Printf("found = %s\n", strconv.FormatBool(found))
	fmt.Printf("result = %d\n", result)
	fmt.Printf("nmap[10] должно равняться \"\", по факту равно \"%s\"\n", nmap[10])

	// мапа со значением nil должна по любому (но соотв типа) ключу отдавать "нулевое" значение подлежащего типа
	var emap map[string]int
	fmt.Printf("emap[vaka] = %d\n", emap["vaka"])

	//получается мапы работают достаточно гибко и вместо ошибок отдают нулевое значение подлежащего типа
	// реальный способ поломаться - попытаться присовить значение неподлежащего типа или обратить с индексом несоотв  типа
	// Важно! Также присвоить что-то nil-ой мапе вызовет поломку
	//emap["data"] = 3 - например вот так
}
