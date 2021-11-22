package main

import "fmt"

func printStringComparison(a string, b string) {
	fmt.Printf("\n%s == %s : %v\n", a, b, a == b)
	fmt.Printf("%s > %s : %v\n", a, b, a > b)
	fmt.Printf("%s < %s : %v\n", a, b, a < b)
}

//Строки можно проверить не тольо на полное равенство, но и на порядок
//В этом случае строку можно представить как число в системе счисления из одних букв алфавита

func main() {
	printStringComparison("mother", "father")
	printStringComparison("a", "b")
	printStringComparison("aa", "ab")
	printStringComparison("vaka", "vaka")
	printStringComparison("str1", "str2")
	printStringComparison("nab", "aaa")
	printStringComparison("Привет", "Привет2")
}
