package parts

import (
	"bufio"
	"sync"

	"adventofcode/internal"
	"adventofcode/pkg/io"
)

func BadgePriorityTotalParallelGroups(input string) int {
	if input == "" {
		input = internal.Input
	}

	file, closeFn := io.OpenInput(input)
	defer closeFn(file)

	var (
		scanner = bufio.NewScanner(file)
		total   int
		wg      sync.WaitGroup
	)

	for {
		rucksacks := make([]string, 0, groupSize)
		for i := 0; i < groupSize; i++ {
			if !scanner.Scan() {
				break
			}

			rucksacks = append(rucksacks, scanner.Text())
		}

		if len(rucksacks) == 0 {
			break
		}

		wg.Add(1)
		go func(group []string) {
			total += findGroupBadgePriority(group)
			wg.Done()
		}(rucksacks)
	}

	wg.Wait()
	return total
}

func BadgePriorityTotalChannels(input string) int {
	if input == "" {
		input = "group.txt"
	}

	file, closeFn := io.OpenInput(input)
	defer closeFn(file)

	var (
		scanner = bufio.NewScanner(file)
		total   int
		wg      sync.WaitGroup
		results = make(chan int)
		groups  = make(chan []string)
	)

	wg.Add(1)
	go func() {
		for {
			group := make([]string, 0, groupSize)
			for i := 0; i < groupSize; i++ {
				if !scanner.Scan() {
					break
				}

				group = append(group, scanner.Text())
			}

			if len(group) == 0 {
				break
			}

			groups <- group
		}

		close(groups)
		wg.Done()
	}()

	for group := range groups {
		wg.Add(1)
		go func(group []string) {
			results <- findGroupBadgePriority(group)
			wg.Done()
		}(group)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		total += result
	}

	return total
}
