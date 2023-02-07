package parts

import (
	"bufio"
	"strconv"

	"adventofcode/internal"
	"adventofcode/pkg/io"
)

func MostCalories(input string) int {
	if input == "" {
		input = internal.Input
	}

	var (
		file, closeFn = io.OpenInput(input)
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
