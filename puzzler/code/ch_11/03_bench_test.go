package main

import (
	"testing"
)

func Multiply(a, b int) int {
	return a * b
}

func BenchmarkMultiply(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Multiply(2, 3)
	}
}

func BenchmarkMultiplyWithSetup(b *testing.B) {
	b.StopTimer()

	a, c := 2, 3

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Multiply(a, c)
	}
}
