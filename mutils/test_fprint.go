package mutils

import (
	"fmt"
	. "octadev.ru/GoLearn/mtypes"
	"os"
)

func TestFprint() {
	mfile, err := os.Create("build/test_fprint.txt")
	if err != nil {
		panic(err)
	}

	drp := 15
	me := Person{Name: "Женя", Surname: "Дронов"}
	bytesWritten, err := fmt.Fprint(mfile, "vaka = ", drp, " which is drp, me = ", me)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Fprint было записано байтов: %d\n", bytesWritten)
}
