package mutils

import "testing"

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
