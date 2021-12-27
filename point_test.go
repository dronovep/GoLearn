package main

import "testing"

func BenchmarkSetCoordsAlg(b *testing.B) {
	p := Point{0, 0, StandardSetCoords}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10; i++ {
			p.SetCoordsAlg(&p, float64(i), 0)
		}
	}
}

func BenchmarkSetCoords(b *testing.B) {
	p := Point{0, 0, StandardSetCoords}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10; i++ {
			p.SetCoords(float64(i), 0)
		}
	}
}
