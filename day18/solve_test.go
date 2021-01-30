package main

import (
	"fmt"
	"testing"
)

func TestSolveP1(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{`1 + 2 * 3 + 4 * 5 + 6`, 71},
		{`1 + (2 * 3) + (4 * (5 + 6))`, 51},
		{`2 * 3 + (4 * 5)`, 26},
		{`5 + (8 * 3 + 9 + 3 * 4 * 3)`, 437},
		{`5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))`, 12240},
		{`((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`, 13632},
	}
	for _, test := range tests {
		testname := fmt.Sprintf("%s", test.input)
		t.Run(testname, func(t *testing.T) {
			ans := solveP1(test.input)
			if ans != test.want {
				t.Errorf("got %d, want %d", ans, test.want)
			}
		})
	}
}
