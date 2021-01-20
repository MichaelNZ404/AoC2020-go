package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestSolveP1(t *testing.T) {
	var tests = []struct {
		input []policy
		want  int
	}{
		{
			[]policy{
				policy{min: 1, max: 3, char: "a", password: "abcde"},
				policy{min: 1, max: 3, char: "b", password: "cdefg"},
				policy{min: 2, max: 9, char: "c", password: "ccccccccc"},
			},
			2},
	}
	for _, test := range tests {
		testname := fmt.Sprintf("%s", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(test.input)), ","), "[]"))
		t.Run(testname, func(t *testing.T) {
			ans := solveP1(test.input)
			if ans != test.want {
				t.Errorf("got %d, want %d", ans, test.want)
			}
		})
	}
}

// func TestSolveP2(t *testing.T) {
// 	var tests = []struct {
// 		input []int
// 		want  int
// 	}{
// 		{[]int{1721, 979, 366, 299, 675, 1456}, 241861950},
// 	}
// 	for _, test := range tests {
// 		testname := fmt.Sprintf("%s", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(test.input)), ","), "[]"))
// 		t.Run(testname, func(t *testing.T) {
// 			ans, _ := solveP2(test.input)
// 			if ans != test.want {
// 				t.Errorf("got %d, want %d", ans, test.want)
// 			}
// 		})
// 	}
// }
