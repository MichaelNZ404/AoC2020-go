package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readInput(filename string) (nums []int, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	nums = make([]int, 0, len(lines))

	for _, line := range lines {
		if len(line) == 0 {
			continue
		} //empty line at EOF
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}

func solveP1(intList []int) (int, error) {
	for indexA, inputA := range intList {
		for indexB, inputB := range intList {
			if inputA+inputB == 2020 && indexA != indexB {
				return inputA * inputB, nil
			}
		}
	}
	return -1, errors.New("Cannot find solution")
}

func main() {
	fmt.Println("Solving Part One!")
	p1Input, err := readInput("input1.txt")
	if err != nil {
		panic(err)
	}
	p1Solution, err := solveP1(p1Input)
	if err != nil {
		panic(err)
	}
	fmt.Println(p1Solution)
}
