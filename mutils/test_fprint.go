package mutils

import (
	"fmt"
	. "octadev.ru/GoLearn/mtypes"
	"os"
)

func TestFprint() {
	mfile, err := os.Create("test_fprint.txt")
	if err != nil {
		panic(err)
	}

	drp := 15
	me := Person{Name: "Женя", Surname: "Дронов"}
	bytes_written, err := fmt.Fprint(mfile, "vaka = ", drp, " which is drp, me = ", me)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Было записано байтов: %d\n", bytes_written)
}
