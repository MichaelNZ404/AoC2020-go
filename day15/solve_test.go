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
		{`0,3,6`, 436},
		{`1,3,2`, 1},
		{`2,1,3`, 10},
		{`1,2,3`, 27},
		{`2,3,1`, 78},
		{`3,2,1`, 438},
		{`3,1,2`, 1836},
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
