package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type policy struct {
	min      int
	max      int
	char     string //change to byte maybe?
	password string
}

func readInput(filename string) (returnStrings []string, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	returnStrings = make([]string, 0, len(lines))

	for _, line := range lines {
		if len(line) == 0 {
			continue
		} //empty line at EOF
		returnStrings = append(returnStrings, line)
	}
	return returnStrings, nil
}

func solveP1(stringList []string) int {
	treesHit := 0
	depth := 0
	distance := 0
	for depth < len(stringList) {
		if stringList[depth][distance] == '#' {
			treesHit = treesHit + 1 //ow!
		}
		distance = (distance + 3) % len(stringList[0])
		depth = depth + 1
	}
	return treesHit
}

// func solveP2(policyList []policy) int {
// 	valid := 0
// 	for _, policy := range policyList {
// 		minExists := string([]byte{policy.password[policy.min-1]}) == policy.char
// 		maxExists := string([]byte{policy.password[policy.max-1]}) == policy.char
// 		if (minExists && !maxExists) || (maxExists && !minExists) {
// 			valid = valid + 1
// 		}
// 	}
// 	return valid
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
	// p2Input, err := readInput("input.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// p2Solution := solveP2(p2Input)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(p2Solution)
}
