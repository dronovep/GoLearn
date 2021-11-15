package main

import "fmt"

func createCounter(init int) func() int {
	start := init

	return func() int {
		out := start
		start += 1
		return out
	}
}

func main() {

	counters := [5]func() int{}

	for i := 0; i < 5; i++ {
		counters[i] = createCounter(i)
	}

	for i := 0; i < 5; i++ {
		for k := 0; k < 3; k++ {
			fmt.Printf("count = %d\n", counters[i]())
		}
		fmt.Printf("next counter\n")
	}

}
