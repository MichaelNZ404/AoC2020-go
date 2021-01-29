package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
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
	values := strings.Split(inputString, ",")
	numberMap := make(map[int]int)
	turnCount := 0 //AKA spoken count
	nextNumberSpoken := 0

	for idx, value := range values {
		number, _ := strconv.Atoi(value)
		if _, exists := numberMap[number]; !exists {
			nextNumberSpoken = 0
		} else {
			nextNumberSpoken = turnCount - numberMap[number]
		}
		numberMap[number] = idx + 1
	}

	turnCount = len(values)
	thisNumberSpoken := 0
	for turnCount < 2020 {
		turnCount++
		thisNumberSpoken = nextNumberSpoken

		if _, exists := numberMap[nextNumberSpoken]; !exists {
			nextNumberSpoken = 0
		} else {
			nextNumberSpoken = turnCount - numberMap[nextNumberSpoken]
		}
		numberMap[thisNumberSpoken] = turnCount
	}
	return thisNumberSpoken
}

func main() {
	fmt.Println("Solving Part One!")
	p1Input := readInput("input.txt")
	p1Solution := solveP1(p1Input)
	fmt.Println(p1Solution)

	// fmt.Println("Solving Part Two!")
	// p2Input := readInput("input.txt")
	// p2Solution := solveP2(p2Input)
	// fmt.Println(p2Solution)
}
