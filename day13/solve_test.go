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
		{
			`939
7,13,x,x,59,x,31,19`, 295},
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

func TestSolveP2(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{
			`939
7,13,x,x,59,x,31,19`, 1068781},
		{
			`939
17,x,13,19`, 3417},
		{
			`939
67,7,59,61`, 754018},
		{
			`939
67,x,7,59,61`, 779210},
		{
			`939
67,7,x,59,61`, 1261476},
		{
			`939
1789,37,47,1889`, 1202161486},
	}
	for _, test := range tests {
		testname := fmt.Sprintf("%s", test.input)
		t.Run(testname, func(t *testing.T) {
			ans := solveP2(test.input)
			if ans != test.want {
				t.Errorf("got %d, want %d", ans, test.want)
			}
		})
	}
}
