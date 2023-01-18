package main

import (
	"bufio"
	"fmt"
	"strconv"

	"AdventOfCode/pkg/io"
)

func main() {
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
			fmt.Println(currentCals, maxCals)
			if currentCals > maxCals {
				maxCals = currentCals
			}
			currentCals = 0
		}
	}

	if currentCals > maxCals {
		fmt.Println(currentCals)
	} else {
		fmt.Println(maxCals)
	}
}
