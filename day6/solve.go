package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readInput(filename string) (returnString string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func solveP1(inputString string) int {
	answerCount := 0
	groups := strings.Split(inputString, "\n\n")
	for _, group := range groups {
		answerMap := make(map[rune]bool)
		members := strings.Split(group, "\n")
		for _, member := range members {
			for _, answer := range member {
				answerMap[answer] = true
			}
		}
		answerCount = answerCount + len(answerMap)
	}
	return answerCount
}

func solveP2(inputString string) int {
	answerCount := 0
	groups := strings.Split(inputString, "\n\n")
	for _, group := range groups {
		answerMap := make(map[rune]int)
		members := strings.Split(group, "\n")
		for _, member := range members {
			for _, answer := range member {
				answerMap[answer] = answerMap[answer] + 1
			}
		}
		for answer := range answerMap {
			if answerMap[answer] == len(members) {
				answerCount = answerCount + 1
			}
		}
	}
	return answerCount
}

func main() {
	fmt.Println("Solving Part One!")
	p1Input := readInput("input.txt")
	p1Solution := solveP1(p1Input)
	fmt.Println(p1Solution)

	fmt.Println("Solving Part Two!")
	p2Input := readInput("input.txt")
	p2Solution := solveP2(p2Input)
	fmt.Println(p2Solution)
}
