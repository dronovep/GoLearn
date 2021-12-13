package mutils

import (
	"fmt"
	"octadev.ru/GoLearn/mtypes"
	"os"
)

func TestFscanf() {
	defer mcatch()
	mfile, err := os.Open("build/test_fscanf.txt")
	if err != nil {
		panic(err.Error())
	}

	var str1, str2, str3, str4 string
	var me mtypes.Person

	itemsScanned, err := fmt.Fscanf(mfile, "%s \nvakagashira \n%s \n%s \n%s\n%v", &str1, &str2, &str3, &str4, &me) // В реале функция должна принимать указатели или что-то, что реализует интерфейс Scanner!
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		panic(err.Error())
	}
	fmt.Printf("Fscan элементов считано: %d\n", itemsScanned)
	fmt.Printf("str1 = %s; str2 = %s; str3 = %s; str4 = %s\n, me = %#v\n", str1, str2, str3, str4, me)
}
