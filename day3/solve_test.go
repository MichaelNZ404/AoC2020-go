package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestSolveP1(t *testing.T) {
	var tests = []struct {
		input []string
		want  int
	}{
		{
			[]string{
				"..##.......",
				"#...#...#..",
				".#....#..#.",
				"..#.#...#.#",
				".#...##..#.",
				"..#.##.....",
				".#.#.#....#",
				".#........#",
				"#.##...#...",
				"#...##....#",
				".#..#...#.#",
			},
			7},
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

func TestSolveP2(t *testing.T) {
	var tests = []struct {
		input []string
		want  int
	}{
		{
			[]string{
				"..##.......",
				"#...#...#..",
				".#....#..#.",
				"..#.#...#.#",
				".#...##..#.",
				"..#.##.....",
				".#.#.#....#",
				".#........#",
				"#.##...#...",
				"#...##....#",
				".#..#...#.#",
			},
			336},
	}
	for _, test := range tests {
		testname := fmt.Sprintf("%s", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(test.input)), ","), "[]"))
		t.Run(testname, func(t *testing.T) {
			ans := solveP2(test.input)
			if ans != test.want {
				t.Errorf("got %d, want %d", ans, test.want)
			}
		})
	}
}
