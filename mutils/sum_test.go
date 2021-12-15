package mutils

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	got := Sum(3, 2, 1, 6, 7)
	if got != 19 {
		t.Errorf("Sum(3, 2, 1, 6, 7) = %d, must be 19\n", got)
	}
}

func ExampleSum() {
	fmt.Printf("Sum(10, 20, 30, 40) = %d", Sum(10, 20, 30, 40))
	// Output:
	// Sum(10, 20, 30, 40) = 100
}
