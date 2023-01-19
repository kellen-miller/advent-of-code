package main

import (
	"fmt"

	"AdventOfCode/internal/2022/day2/parts"
)

func main() {
	fmt.Println("--- Day 2: Rock Paper Scissors ---")

	fmt.Println("Part 1:", "Total score =>", parts.TotalScore())

	fmt.Println("Part 2:", "Set round result =>", parts.SetRoundResult())
}
