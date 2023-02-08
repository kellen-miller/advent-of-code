package parts

import (
	"strconv"
	"strings"

	"adventofcode/internal"
	"adventofcode/pkg/io"
)

func RedundantCleanup(input string) int {
	if input == "" {
		input = internal.Input
	}

	sc, closeFile := io.GetScanner(input)
	defer closeFile()

	var total int
	for sc.Scan() {
		total += isAssignmentRedundant(sc.Text())
	}
	return total
}

func isAssignmentRedundant(pairs string) int {
	elves := strings.Split(pairs, ",")

	elf1Min, elf1Max := getElfSections(elves[0])
	elf2Min, elf2Max := getElfSections(elves[1])

	if (elf1Min <= elf2Min && elf1Max >= elf2Max) ||
		(elf2Min <= elf1Min && elf2Max >= elf1Max) {
		return 1
	}

	return 0
}

func getElfSections(elf string) (int, int) {
	sections := strings.Split(elf, "-")
	section1, _ := strconv.Atoi(sections[0])
	section2, _ := strconv.Atoi(sections[1])
	return section1, section2
}
