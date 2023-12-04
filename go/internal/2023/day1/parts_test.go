package day1

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tcs := map[string]struct {
		input string
		want  int
	}{
		"example": {
			input: "data/example1.txt",
			want:  142,
		},
		"input": {
			input: "data/input.txt",
			want:  54630,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			got := Calibration(tc.input)
			if got != tc.want && tc.want != 0 {
				t.Errorf("Calibration = %d; want %d", got, tc.want)
			} else {
				t.Logf("Calibration = %d", got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tcs := map[string]struct {
		input string
		want  int
	}{
		"example": {
			input: "data/example2.txt",
			want:  281,
		},
		"input": {
			input: "data/input.txt",
			want:  54770,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			got := CalibrationSpelled(tc.input)
			if got != tc.want && tc.want != 0 {
				t.Errorf("Calibration = %d; want %d", got, tc.want)
			} else {
				t.Logf("Calibration = %d", got)
			}
		})
	}
}
