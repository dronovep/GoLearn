package main

import "fmt"

func printWords(start string, words ...string) {
	output := ""
	for i := range words {
		output += words[i] + " "
	}

	fmt.Printf("%s: %s\n", start, output)
}

func dumpArgs(args ...string) {
	fmt.Printf("%v\n", args)
}

func dumpArgs2(args ...string) {
	if args != nil {
		fmt.Printf("%v\n", args)
	} else {
		fmt.Printf("nil\n")
	}
}

func processInts(args ...int) {
	for i := range args {
		args[i] = i
	}
}

var mis = [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

func main() {
	printWords("Words are", "Mother", "Father", "Sister", "Brother")
	dumpArgs("Kia", "Sorento", "Audi", "Q7", "Chevrolet", "Tahoe")
	dumpArgs()
	dumpArgs2("Mother", "Father", "Sister", "Brother")
	dumpArgs2()

	fmt.Printf("mis до вызова processInts на его фрагменте: %v\n", mis)
	//processInts(mis...)   массив в качестве слайса использовать не получится
	processInts(mis[2:8]...)
	fmt.Printf("mis после вызова processInts на его фрагменте: %v\n", mis)
}
