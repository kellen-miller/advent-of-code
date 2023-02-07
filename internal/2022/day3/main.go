package main

import (
	"fmt"

	"adventofcode/internal/2022/day3/parts"
)

func main() {
	input := "internal/2022/day3/input.txt"

	fmt.Println("--- Day 3: Rucksack Reorganization ---")

	fmt.Println("Part 1:", "Total priority =>", parts.SuppliesPriorityTotal(input))

	// fmt.Println("Part 2:", "Set round result =>", parts.SetRoundResult())
}
