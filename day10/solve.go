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

// func getSubtreeCount(currentVal int, targetVal int, intList []int) int {
// 	count := 0
// 	for idx, iVal := range intList {
// 		if iVal-currentVal > 3 {
// 			break
// 		} else {
// 			count = count + getSubtreeCount(iVal, targetVal, intList[idx+1:])
// 		}
// 		if targetVal-iVal < 4 {
// 			count = count + 1
// 		}
// 	}
// 	return count
// }

// This solution has bad performance
// func solveP2(inputString string) int {
// 	stringList := strings.Split(inputString, "\n")
// 	intList := make([]int, len(stringList))
// 	for sIndex, sInput := range stringList {
// 		newInt, err := strconv.Atoi(sInput)
// 		if err != nil {
// 			panic(err)
// 		}
// 		intList[sIndex] = newInt
// 	}
// 	sort.Ints(intList)

// 	currentVal := 0
// 	targetVal := intList[len(intList)-1] + 3

// 	return getSubtreeCount(currentVal, targetVal, intList)
// }

func getSubtreeCount(currentVal int, targetVal int, intList []int) int {
	if len(intList) == 1 {
		return 1
	}

	count := 0
	for idx := 1; idx < len(intList); idx++ {
		iVal := intList[idx]
		if iVal-currentVal > 3 {
			break
		}
		if iVal == targetVal {
			count = count + 1
		} else {
			count = count + getSubtreeCount(iVal, targetVal, intList[idx:])
		}
	}
	return count
}

func solveP2(inputString string) int {
	stringList := strings.Split(inputString, "\n")
	intList := make([]int, len(stringList))
	for sIndex, sInput := range stringList {
		newInt, err := strconv.Atoi(sInput)
		if err != nil {
			panic(err)
		}
		intList[sIndex] = newInt
	}
	intList = append(intList, 0)
	sort.Ints(intList)
	intList = append(intList, intList[len(intList)-1]+3)
	count := 1

	lastSplit := 0
	subproblems := make([][]int, 0)
	for idx, iVal := range intList {
		if idx+1 > len(intList)-1 {
			subproblems = append(subproblems, intList[lastSplit:])
			break
		}
		if iVal+3 == intList[idx+1] {
			subproblems = append(subproblems, intList[lastSplit:idx+1])
			lastSplit = idx + 1
		}
	}

	for _, subproblem := range subproblems {
		subResult := getSubtreeCount(subproblem[0], subproblem[len(subproblem)-1], subproblem)
		// fmt.Println(subproblem, subResult)
		count = count * subResult
	}
	return count
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
