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

func Benchmark_Day3_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parts.BadgePriorityTotal("input.txt")
	}
}

func Benchmark_Day3_Part2_ParallelGroups(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parts.BadgePriorityTotalParallelGroups("input.txt")
	}
}

func Benchmark_Day3_Part2_Channels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parts.BadgePriorityTotalChannels("input.txt")
	}
}
