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
			`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`, 165},
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
			`mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`, 208},
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
