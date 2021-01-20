package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type policy struct {
	min      int
	max      int
	char     string //change to byte maybe?
	password string
}

func readInput(filename string) (returnPolicies []policy, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	returnPolicies = make([]policy, 0, len(lines))

	for _, line := range lines {
		if len(line) == 0 {
			continue
		} //empty line at EOF

		returnPolicy := policy{}
		//line of form "1-3 a: abcde"
		chunks := strings.Fields(line)
		if len(chunks) != 3 {
			return nil, errors.New("Unexpected input string")
		}
		minmax := strings.Split(chunks[0], "-")
		min, err := strconv.Atoi(minmax[0])
		if err != nil {
			return nil, err
		}
		returnPolicy.min = min
		max, err := strconv.Atoi(minmax[1])
		if err != nil {
			return nil, err
		}
		returnPolicy.max = max
		char := strings.Trim(chunks[1], ":")
		returnPolicy.char = char
		returnPolicy.password = chunks[2]
		returnPolicies = append(returnPolicies, returnPolicy)
	}
	return returnPolicies, nil
}

func solveP1(policyList []policy) int {
	valid := 0
	for _, policy := range policyList {
		policyCount := strings.Count(policy.password, policy.char)
		if policyCount >= policy.min && policyCount <= policy.max {
			valid = valid + 1
		}
	}
	return valid
}

// func solveP2(intList []int) (int, error) {
// 	for indexA, inputA := range intList {
// 		for indexB, inputB := range intList {
// 			for indexC, inputC := range intList {
// 				if inputA+inputB+inputC == 2020 && indexA != indexB && indexA != indexC && indexC != indexB {
// 					return inputA * inputB * inputC, nil
// 				}
// 			}
// 		}
// 	}
// 	return -1, errors.New("Cannot find solution")
// }

func main() {
	fmt.Println("Solving Part One!")
	p1Input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}
	p1Solution := solveP1(p1Input)
	fmt.Println(p1Solution)

	// fmt.Println("Solving Part Two!")
	// p2Input, err := readInput("input1.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// p2Solution, err := solveP2(p2Input)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(p2Solution)
}
