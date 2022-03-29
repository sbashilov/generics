package main

// go test -bench . -benchmem

import (
	"testing"
)

func BenchmarkNoGenerics(b *testing.B) {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}
	for i := 0; i < b.N; i++ {
		SumInts(ints)
		SumFloats(floats)

	}
}

func BenchmarkGenerics(b *testing.B) {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}
	for i := 0; i < b.N; i++ {
		SumIntsOrFloats(ints)
		SumIntsOrFloats(floats)

	}
}
