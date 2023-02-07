package main

import (
	"testing"

	"adventofcode/internal/2022/day2/parts"
)

func Benchmark_Day2_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parts.TotalScore("input.txt")
	}
}

func Benchmark_Day2_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parts.SetRoundResult("input.txt")
	}
}
