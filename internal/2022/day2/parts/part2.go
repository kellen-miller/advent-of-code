package parts

import (
	"bufio"

	"adventofcode/internal"
	"adventofcode/pkg/io"
)

const (
	choices = 3
)

func SetRoundResult(input string) int {
	if input == "" {
		input = internal.Input
	}

	file, closeFn := io.OpenInput(input)
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
			total += (oppChoice + 1) % choices
		case draw:
			total += oppChoice
		case loss:
			total += (oppChoice + 2) % choices
		}
	}

	return total
}
