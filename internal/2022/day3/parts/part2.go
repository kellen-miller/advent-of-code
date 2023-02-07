package parts

import (
	"bufio"

	"adventofcode/internal"
	"adventofcode/pkg/io"
)

const (
	groupSize = 3
)

func BadgePriorityTotal(input string) int {
	if input == "" {
		input = internal.Input
	}

	file, closeFn := io.OpenInput(input)
	defer closeFn(file)

	var (
		scanner = bufio.NewScanner(file)
		total   int
	)

	for {
		groups := make([]string, groupSize)
		for i := 0; i < groupSize; i++ {
			if !scanner.Scan() {
				return total
			}

			groups[i] = scanner.Text()
		}

		total += findGroupBadgePriority(groups)
	}
}

func findGroupBadgePriority(groups []string) int {
	countGroupsBadgeFound := make([]int, maxLetterDifference)

	for _, group := range groups {
		badgesFoundInGroup := make([]bool, maxLetterDifference)

		for _, badge := range group {
			badgeVal := badge - 'A'

			if !badgesFoundInGroup[badgeVal] {
				badgesFoundInGroup[badgeVal] = true
				countGroupsBadgeFound[badgeVal]++
			}

			if countGroupsBadgeFound[badgeVal] == len(groups) {
				return letterScore(badge)
			}
		}
	}

	return 0
}
