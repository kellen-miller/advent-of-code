package main

import (
	"testing"

	"adventofcode/internal/2022/day3/parts"
)

func Benchmark_Day3_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parts.SuppliesPriorityTotal("input.txt")
	}
}

// func Benchmark_Day3_Part2(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		parts.SetRoundResult("input.txt")
// 	}
// }
