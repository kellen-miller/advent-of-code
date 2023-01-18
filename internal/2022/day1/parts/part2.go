package parts

import (
	"bufio"
	"strconv"

	"AdventOfCode/pkg/io"

	"github.com/emirpasic/gods/trees/binaryheap"
)

type CapacityHeap struct {
	Capacity int
	*binaryheap.Heap
}

func Top3Calories() ([]interface{}, int) {
	var (
		file, closeFn = io.OpenInput("internal/2022/day1/input.txt")
		sc            = bufio.NewScanner(file)
		elfHeap       = &CapacityHeap{
			Capacity: 3,
			Heap:     binaryheap.NewWithIntComparator(),
		}
		currentCals int
	)
	defer closeFn(file)

	for sc.Scan() {
		cals, err := strconv.Atoi(sc.Text())
		currentCals += cals

		if err != nil {
			elfHeap.push(currentCals)
			currentCals = 0
		}
	}

	var sum int
	for _, val := range elfHeap.Values() {
		sum += val.(int)
	}

	return elfHeap.Values(), sum
}

func (h *CapacityHeap) push(val int) {
	if h.Size() < h.Capacity {
		h.Push(val)
	} else {
		peek, notEmpty := h.Peek()
		if notEmpty && peek.(int) < val {
			h.Pop()
			h.Push(val)
		}
	}
}
