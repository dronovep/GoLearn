package mutils

func Sum(args ...int) int {
	sum := 0
	for _, value := range args {
		sum += value
	}

	return sum
}
