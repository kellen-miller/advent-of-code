package parts

import (
	"bufio"
	"strconv"

	"AdventOfCode/pkg/io"
)

func MostCalories() int {
	var (
		file, closeFn = io.OpenInput("internal/2022/day1/input.txt")
		maxCals       int
		currentCals   int
	)
	defer closeFn(file)

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		cals, err := strconv.Atoi(sc.Text())
		currentCals += cals

		if err != nil {
			if currentCals > maxCals {
				maxCals = currentCals
			}
			currentCals = 0
		}
	}

	if currentCals > maxCals {
		maxCals = currentCals
	}

	return maxCals
}
