package mutils

import (
	"sync"
	"testing"
)

func TestMul(t *testing.T) {
	t.Run("first", func(t *testing.T) {
		got := Mul(3, 5, 6)
		if got != 90 {
			t.Errorf("Mul(3, 5, 6) = %d, want 90", got)
		}
	})

	t.Run("second", func(t *testing.T) {
		got := Mul(1, 2, 3)
		if got != 6 {
			t.Errorf("Mul(1, 2, 3) = %d, want 6", got)
		}
	})

	t.Run("third", func(t *testing.T) {
		got := Mul(3, 4, 6)
		if got != 72 {
			t.Errorf("Mul(3, 4, 6) = %d, want 72", got)
		}
	})
}

func tstfirst(t *testing.T) {
	got := Mul(3, 5, 6)
	if got != 90 {
		t.Errorf("Mul(3, 5, 6) = %d, want 90", got)
	}
}

func tstsecond(t *testing.T) {
	got := Mul(1, 2, 3)
	if got != 6 {
		t.Errorf("Mul(1, 2, 3) = %d, want 6", got)
	}
}

func tstthird(t *testing.T) {
	got := Mul(3, 4, 6)
	if got != 72 {
		t.Errorf("Mul(3, 4, 6) = %d, want 72", got)
	}
}

func TestMul2(t *testing.T) {
	t.Run("first", tstfirst)
	t.Run("second", tstsecond)
	t.Run("third", tstthird)
}

// По умолчанию go test выполнит напрямую все тесты
// Но если хочется, чтобы тест пакета выполнялся как-то по особенному, можно сделать в пакете функцию TestMain(m *testing.M)
// например для вынесения setups и teardowns
// TestMain всегда выполняется в main goroutine
func TestMain(m *testing.M) {
	m.Run()
}

func AsyncAssign() int {
	value := 0
	for i := 0; i <= 100; i++ {
		value = i
	}

	return value
}

func SyncAssign(ch chan bool) int {
	value := 0
	for i := 0; i <= 100; i++ {
		ch <- true
		value = i
		<-ch
	}

	return value
}

func MtxAssign() int {
	mtx := sync.Mutex{}
	value := 0
	for i := 0; i <= 100; i++ {
		mtx.Lock()
		value = i
		mtx.Unlock()
	}

	return value
}

func BenchmarkAsyncAssign(t *testing.B) {
	for i := 0; i < t.N; i++ {
		AsyncAssign()
	}
}

func BenchmarkSyncAssign(t *testing.B) {
	ch := make(chan bool, 1)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		SyncAssign(ch)
	}
}

func BenchmarkMtxAssign(t *testing.B) {
	for i := 0; i < t.N; i++ {
		MtxAssign()
	}
}
