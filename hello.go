package main

import (
	"fmt"
	. "octadev.ru/GoLearn/mutils"
)

func main() {
	fmt.Printf("\nНачало работы!\n\n")
	TestPrintfVformats()
	TestErrorf()
	TestFprint()
	TestFprintF()
	TestFscan()
	TestFscanf()
	fmt.Printf("\nРабота завершена!\n")
}
