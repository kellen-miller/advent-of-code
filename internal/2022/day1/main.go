package main

import (
	"fmt"

	"AdventOfCode/internal/2022/day1/parts"
)

func main() {
	fmt.Println("--- Day 1: Calorie Counting ---")

	fmt.Println("Part 1:", "Most calories being carried =>", parts.MostCalories())

	vals, sum := parts.Top3Calories()
	fmt.Println("Part 2:", fmt.Sprintf("Sum of top 3 calories carried = sum(%v) => %d\n", vals, sum))
}
