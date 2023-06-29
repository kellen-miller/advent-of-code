package parts

import (
	"github.com/kellen-miller/advent-of-code/go/internal/common"
	"github.com/kellen-miller/advent-of-code/go/pkg/io"
)

const (
	messageStartSize = 14
)

func StartOfMessage(input string) []int {
	if input == "" {
		input = common.Input
	}

	sc, closeFn := io.GetScanner(input)
	defer closeFn()

	var ms []int
	for sc.Scan() {
		ms = append(ms, findUniqueSetOfSize(sc.Text(), messageStartSize))
	}

	return ms
}
