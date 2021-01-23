package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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
	stringList := strings.Split(inputString, "\n")
	intList := make([]int, len(stringList))
	for sIndex, sInput := range stringList {
		newInt, err := strconv.Atoi(sInput)
		if err != nil {
			panic(err)
		}
		intList[sIndex] = newInt
	}
	sort.Ints(intList)

	oneDiff := 0
	threeDiff := 1 // final extra is always +3

	for i, val := range intList {
		diff := 0
		if i == 0 {
			diff = val
		} else {
			diff = val - intList[i-1]
		}

		if diff == 1 {
			oneDiff = oneDiff + 1
		}
		if diff == 3 {
			threeDiff = threeDiff + 1
		}
	}
	return oneDiff * threeDiff
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
