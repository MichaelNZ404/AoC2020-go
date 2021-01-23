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
			`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`, 127},
	}
	for _, test := range tests {
		testname := fmt.Sprintf("%s", test.input)
		t.Run(testname, func(t *testing.T) {
			ans := solveP1(test.input, 5)
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
			`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`, 62},
	}
	for _, test := range tests {
		testname := fmt.Sprintf("%s", test.input)
		t.Run(testname, func(t *testing.T) {
			ans := solveP2(test.input, 5)
			if ans != test.want {
				t.Errorf("got %d, want %d", ans, test.want)
			}
		})
	}
}
