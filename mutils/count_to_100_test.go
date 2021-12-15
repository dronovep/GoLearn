package mutils

import (
	"fmt"
	"testing"
)

func BenchmarkCountTo100(b *testing.B) {
	fmt.Printf("b.N = %d\n", b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountTo100()
	}
}
