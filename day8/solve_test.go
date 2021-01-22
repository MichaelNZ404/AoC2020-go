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
			`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`, 5},
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

// func TestSolveP2(t *testing.T) {
// 	var tests = []struct {
// 		input string
// 		want  int
// 	}{
// 		{
// 			`shiny gold bags contain 2 dark red bags.
// dark red bags contain 2 dark orange bags.
// dark orange bags contain 2 dark yellow bags.
// dark yellow bags contain 2 dark green bags.
// dark green bags contain 2 dark blue bags.
// dark blue bags contain 2 dark violet bags.
// dark violet bags contain no other bags.`, 126},
// 	}
// 	for _, test := range tests {
// 		testname := fmt.Sprintf("%s", test.input)
// 		t.Run(testname, func(t *testing.T) {
// 			ans := solveP2(test.input)
// 			if ans != test.want {
// 				t.Errorf("got %d, want %d", ans, test.want)
// 			}
// 		})
// 	}
// }
