package mutils

func Mul(args ...int) int {
	mul := 1
	for _, value := range args {
		mul *= value
	}

	return mul
}
