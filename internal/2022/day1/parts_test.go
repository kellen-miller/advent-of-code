package main

import (
	"testing"

	"adventofcode/internal/2022/day1/parts"
)

func Benchmark_Day1_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parts.MostCalories("input.txt")
	}
}

func Benchmark_Day1_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parts.Top3Calories("input.txt")
	}
}
