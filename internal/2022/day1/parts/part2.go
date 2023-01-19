package parts

import (
	"bufio"
	"strconv"

	"AdventOfCode/pkg/io"
	"AdventOfCode/pkg/structures"

	"github.com/ugurcsen/gods-generic/trees/binaryheap"
)

func Top3Calories() ([]int, int) {
	var (
		file, closeFn = io.OpenInput("internal/2022/day1/input.txt")
		sc            = bufio.NewScanner(file)
		elfHeap       = &structures.CapacityHeap[int]{
			Capacity: 3,
			Heap:     binaryheap.NewWithNumberComparator[int](),
		}
		currentCals int
	)
	defer closeFn(file)

	for sc.Scan() {
		cals, err := strconv.Atoi(sc.Text())
		currentCals += cals

		if err != nil {
			elfHeap.Push(currentCals)
			currentCals = 0
		}
	}

	var sum int
	for _, val := range elfHeap.Values() {
		sum += val
	}

	return elfHeap.Values(), sum
}
