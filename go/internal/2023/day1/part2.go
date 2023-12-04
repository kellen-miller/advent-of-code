package day1

import (
	"log"
	"strconv"
	"unicode"

	"github.com/kellen-miller/advent-of-code/go/pkg/io"
)

var numWordLenToNum = map[int][]string{
	3: {"one", "two", "six"},
	4: {"zero", "four", "five", "nine"},
	5: {"three", "seven", "eight"},
}

var numWordToNum = map[string]uint8{
	"zero":  '0',
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func CalibrationSpelled(input string) int {
	sc, closeFile := io.GetScanner(input)
	defer closeFile()

	var sum int

	for sc.Scan() {
		line := sc.Text()

		left := findLeftChar(line)
		right := findRightChar(line)

		val, err := strconv.Atoi(string(left) + string(right))
		if err != nil {
			log.Panic(err)
		}

		sum += val
	}

	return sum
}

func findLeftChar(line string) uint8 {
	for i := 0; i < len(line); i++ {
		if unicode.IsDigit(rune(line[i])) {
			return line[i]
		}

		for k, v := range numWordLenToNum {
			if len(line)-i < k {
				continue
			}

			for _, word := range v {
				if line[i:i+k] == word {
					return numWordToNum[word]
				}
			}
		}
	}

	return 0
}

func findRightChar(line string) uint8 {
	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			return line[i]
		}

		for k, v := range numWordLenToNum {
			if i < k {
				continue
			}

			for _, word := range v {
				if line[i-k+1:i+1] == word {
					return numWordToNum[word]
				}
			}
		}
	}

	return 0
}
