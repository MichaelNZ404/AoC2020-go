package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
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

type ruleRange struct {
	min int
	max int
}

func solveP1(inputString string) int {
	sections := strings.Split(inputString, "\n\n")
	rules := strings.Split(sections[0], "\n")
	bounds := make([]ruleRange, 0)

	for _, rule := range rules {
		r, _ := regexp.Compile(`(\d+)-(\d+)`)
		digits := r.FindAllStringSubmatch(rule, -1)
		for _, submatch := range digits {
			min, _ := strconv.Atoi(submatch[1])
			max, _ := strconv.Atoi(submatch[2])
			bounds = append(bounds, ruleRange{min: min, max: max})
		}
	}

	nearbyTickets := strings.Split(sections[2], "\n")
	errorSum := 0
	for _, ticket := range nearbyTickets {
		fields := strings.Split(ticket, ",")
		for _, field := range fields {
			intField, _ := strconv.Atoi(field)

			valid := false
			for _, bound := range bounds {
				if intField <= bound.max && intField >= bound.min {
					valid = true
					break
				}
			}
			if valid == false {
				errorSum = errorSum + intField
			}
		}
	}
	return errorSum
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
