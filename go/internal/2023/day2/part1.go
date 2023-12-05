package day2

import (
	"strconv"
	"strings"

	"github.com/kellen-miller/advent-of-code/go/pkg/io"
)

const (
	MaxRedCubes   = 12
	MaxGreenCubes = 13
	MaxBlueCubes  = 14
)

func GamesPossibleSum(input string) int {
	sc, closeFile := io.GetScanner(input)
	defer closeFile()

	var sum int

	for sc.Scan() {
		line := sc.Text()

		id, possible := parseGame(line)
		if possible {
			sum += id
		}
	}

	return sum
}

func parseGame(line string) (int, bool) {
	const gamePartsWant = 2

	gameParts := strings.Split(line, ":")
	if len(gameParts) != gamePartsWant {
		panic("invalid game")
	}

	return parseID(gameParts[0]), parseRounds(gameParts[1])
}

func parseID(gamePart string) int {
	const idPartsWant = 2

	idParts := strings.Split(gamePart, " ")
	if len(idParts) != idPartsWant {
		panic("invalid game id")
	}

	id, err := strconv.Atoi(idParts[1])
	if err != nil {
		panic("invalid game id")
	}

	return id
}

func parseRounds(roundPart string) bool {
	for _, round := range strings.Split(roundPart, ";") {
		for _, cube := range strings.Split(round, ",") {
			if !parseCube(cube) {
				return false
			}
		}
	}

	return true
}

func parseCube(cube string) bool {
	const cubePartsWant = 2

	cubeParts := strings.Split(strings.TrimSpace(cube), " ")
	if len(cubeParts) != cubePartsWant {
		panic("invalid cube part")
	}

	cubes, err := strconv.Atoi(cubeParts[0])
	if err != nil {
		panic("invalid cube amount")
	}

	switch cubeParts[1] {
	case "red":
		if cubes > MaxRedCubes {
			return false
		}
	case "green":
		if cubes > MaxGreenCubes {
			return false
		}
	case "blue":
		if cubes > MaxBlueCubes {
			return false
		}
	default:
		panic("unknown cube color")
	}

	return true
}
