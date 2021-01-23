package main

import (
	"errors"
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

func solveP1(inputString string, preambleLength int) int {
	stringList := strings.Split(inputString, "\n")
	intList := make([]int, len(stringList))
	for sIndex, sInput := range stringList {
		newInt, err := strconv.Atoi(sInput)
		if err != nil {
			panic(err)
		}
		intList[sIndex] = newInt
	}

	currentLoc := preambleLength
	for currentLoc < len(intList) {
		intsToCheck := intList[currentLoc-preambleLength : currentLoc]
		solution := false
		for i, iValue := range intsToCheck {
			if solution {
				break
			}
			for j, jValue := range intsToCheck {
				if i != j && iValue+jValue == intList[currentLoc] {
					solution = true
					break
				}
			}
		}
		if !solution {
			return intList[currentLoc]
		}
		currentLoc = currentLoc + 1
	}
	panic(errors.New("We made it through"))
}

func solveP2(inputString string, preambleLength int) int {
	stringList := strings.Split(inputString, "\n")
	intList := make([]int, len(stringList))
	for sIndex, sInput := range stringList {
		newInt, err := strconv.Atoi(sInput)
		if err != nil {
			panic(err)
		}
		intList[sIndex] = newInt
	}

	// get the p1 solution
	targetValue := -1
	currentLoc := preambleLength
	for currentLoc < len(intList) {
		if targetValue > 0 {
			break
		}
		intsToCheck := intList[currentLoc-preambleLength : currentLoc]
		solution := false
		for i, iValue := range intsToCheck {
			if solution {
				break
			}
			for j, jValue := range intsToCheck {
				if i != j && iValue+jValue == intList[currentLoc] {
					solution = true
					break
				}
			}
		}
		if !solution {
			targetValue = intList[currentLoc]
		}
		currentLoc = currentLoc + 1
	}

	// find the solution range
	rangeLower := -1
	rangeUpper := -1
	for i, iValue := range intList {
		if rangeLower > 0 || rangeUpper > 0 {
			break
		}
		thisISum := iValue
		for j, jValue := range intList[i+1:] {
			thisISum = thisISum + jValue
			if thisISum == targetValue {
				rangeLower = i
				rangeUpper = j + i + 2
				break
			}
			if thisISum > targetValue {
				break
			}
		}
	}

	// find the min & max of this range
	rangeMin := intList[rangeLower]
	rangeMax := intList[rangeLower]
	for _, val := range intList[rangeLower:rangeUpper] {
		if val > rangeMax {
			rangeMax = val
		}
		if val < rangeMin {
			rangeMin = val
		}
	}
	return rangeMin + rangeMax
}

func main() {
	fmt.Println("Solving Part One!")
	p1Input := readInput("input.txt")
	p1Solution := solveP1(p1Input, 25)
	fmt.Println(p1Solution)

	fmt.Println("Solving Part Two!")
	p2Input := readInput("input.txt")
	p2Solution := solveP2(p2Input, 25)
	fmt.Println(p2Solution)
}
