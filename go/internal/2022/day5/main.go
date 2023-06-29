package main

import (
	"fmt"

	"adventofcode/internal/2022/day5/parts"
)

func main() {
	input := "internal/2022/day5/input.txt"
	// input := "internal/2022/day5/example.txt"

	fmt.Println("--- Day 5: Supply Stacks ---")

	// fmt.Println("Part 1:", "Rearrange crates =>", parts.RearrangeCrates(input))

	fmt.Println("Part 2:", "Rearrange multiple crates =>", parts.RearrangeCratesMulti(input))
}
