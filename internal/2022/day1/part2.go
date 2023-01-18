package main

import (
	"bufio"
	"fmt"
	"strconv"

	"AdventOfCode/pkg/io"

	"github.com/emirpasic/gods/trees/binaryheap"
)

type HeapWithCapacity struct {
	capacity int
	*binaryheap.Heap
}

func main() {
	var (
		file, closeFn = io.OpenInput("internal/2022/day1/input.txt")
		sc            = bufio.NewScanner(file)
		elfHeap       = &HeapWithCapacity{
			capacity: 3,
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

	fmt.Println("vals: ", elfHeap.Values())
	fmt.Println("sum: ", sum)
}

func (h *HeapWithCapacity) push(val int) {
	if h.Size() < h.capacity {
		h.Push(val)
	} else {
		peek, _ := h.Peek()
		if peek.(int) < val {
			h.Pop()
			h.Push(val)
		}
	}
}
