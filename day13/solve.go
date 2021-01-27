package main

import (
	"fmt"
	"io/ioutil"
	"math"
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
	lines := strings.Split(inputString, "\n")
	earliestTime, _ := strconv.Atoi(lines[0])

	minTime := 0
	minTimeBusID := 0
	for _, val := range strings.Split(lines[1], ",") {
		if val != "x" {
			busID, _ := strconv.Atoi(val)
			candidateTime := int(math.Ceil(float64(earliestTime)/float64(busID))) * busID
			fmt.Println(candidateTime)
			if minTimeBusID == 0 || candidateTime < minTime {
				minTime = candidateTime
				minTimeBusID = busID
			}
		}
	}
	return minTimeBusID * (minTime - earliestTime)
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
