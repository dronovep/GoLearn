package mutils

import (
	"fmt"
	"os"
)

var mvar float64 = 35.56

func TestFprintF() {
	defer mcatch()

	mfile, err := os.Create("build/test_fprintf.txt")
	if err != nil {
		panic(err)
	}
	bytesWritten, err := fmt.Fprintf(mfile, "Пишем в файл переменную a = %.2f", mvar)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Fprintf было записано байтов: %d\n", bytesWritten)
}
