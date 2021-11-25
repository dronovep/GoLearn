package main

import (
	"fmt"
	"os"
)

type rstr func() string

func greet() string {
	return "Calling greet"
}

func bye() string {
	return "Calling bye"
}

func generate_bm() (brand string, model string) {
	//func generate_bm() (string, string) {
	brand = "Kia"
	model = "Sorento"

	//return
	//return brand, model
	return "Volkswagen", "Touareq"
}

func printStringArgs(args ...string) {
	fmt.Printf("\n\nString args are:\n")
	for i := range args {
		fmt.Printf("  %s\n", args[i])
	}
}

var actions []rstr = []rstr{greet, bye}

func sum(a int, b int) (result int) {
	result = a + b
	return 25
}

func main() {

	for i := 0; i < len(actions); i++ {
		fmt.Printf("%s\n", actions[i]())
	}

	a, b := generate_bm()

	fmt.Printf("sum(3, 2) = %d\n", sum(3, 2))

	fmt.Printf("Brand = %s, model = %s\n", a, b)

	printStringArgs("mother", "father", "sister", "brother")

	fmt.Printf("Command arguments are:\n")
	for i := range os.Args {
		fmt.Printf("    %s\n", os.Args[i])
	}
}
