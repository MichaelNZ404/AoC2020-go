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
		{`0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`, 2},
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
