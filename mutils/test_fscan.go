package mutils

import (
	"fmt"
	"os"
)

func TestFscan() {
	defer mcatch()
	mfile, err := os.Open("build/test_fscan.txt")
	if err != nil {
		panic(err.Error())
	}

	var str1, str2, str3, str4 string
	var fl1 float64
	itemsScanned, err := fmt.Fscan(mfile, &str1, &str2, &str3, &str4, &fl1) // В реале функция должна принимать указатели! Хотя в документации тут стоит просто интерфейс
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		panic(err.Error())
	}
	fmt.Printf("Fscan элементов считано: %d\n", itemsScanned)
	fmt.Printf("str1 = %s; str2 = %s; str3 = %s; str4 = %s; fl1 = %.2f\n", str1, str2, str3, str4, fl1)
}
