package parts

import (
	"bufio"

	"AdventOfCode/pkg/io"
)

func SetRoundResult() int {
	file, closeFn := io.OpenInput("internal/2022/day2/input.txt")
	defer closeFn(file)

	var (
		total    int
		scanner  = bufio.NewScanner(file)
		scoreMap = map[uint8]int{
			'A': rock,
			'B': paper,
			'C': scissors,
			'X': loss,
			'Y': draw,
			'Z': win,
		}
	)

	for scanner.Scan() {
		var (
			line         = scanner.Text()
			oppChoice    = scoreMap[line[0]]
			resultWanted = scoreMap[line[2]]
		)

		total += resultWanted + 1
		switch resultWanted {
		case win:
			total += (oppChoice + 1) % 3
		case draw:
			total += oppChoice
		case loss:
			total += (oppChoice + 2) % 3
		}
	}

	return total
}
