package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

func checkSlope(stringList []string, slopeX int, slopeY int) int {
	treesHit := 0
	depth := 0
	distance := 0
	for depth < len(stringList) {
		if stringList[depth][distance] == '#' {
			treesHit = treesHit + 1 //ow!
		}
		distance = (distance + slopeX) % len(stringList[0])
		depth = depth + slopeY
	}
	return treesHit
}

func solveP2(stringList []string) int {
	totalHit := 1
	totalHit = totalHit * checkSlope(stringList, 1, 1)
	totalHit = totalHit * checkSlope(stringList, 3, 1)
	totalHit = totalHit * checkSlope(stringList, 5, 1)
	totalHit = totalHit * checkSlope(stringList, 7, 1)
	totalHit = totalHit * checkSlope(stringList, 1, 2)
	return totalHit
}

func main() {
	fmt.Println("Solving Part One!")
	p1Input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}
	p1Solution := solveP1(p1Input)
	fmt.Println(p1Solution)

	fmt.Println("Solving Part Two!")
	p2Input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}
	p2Solution := solveP2(p2Input)
	if err != nil {
		panic(err)
	}
	fmt.Println(p2Solution)
}
