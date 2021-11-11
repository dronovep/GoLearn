package main

import "fmt"

func main() {

	mmap := make(map[string]int, 4)

	mmap["mother"] = 1
	mmap["father"] = 2
	mmap["brother"] = 3
	mmap["me"] = 4
	mmap["extra"] = 10 //это уже пятый элемент, но расширить видимо можно

	for key := range mmap {
		fmt.Printf("mmap[%s] = %d\n", key, mmap[key])
	}

}
