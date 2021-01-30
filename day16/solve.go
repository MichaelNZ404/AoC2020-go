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

type ticketField struct {
	name      string
	rangeOne  ruleRange
	rangeTwo  ruleRange
	positions []int
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

func solveP2(inputString string) int {
	sections := strings.Split(inputString, "\n\n")
	rules := strings.Split(sections[0], "\n")
	ticketFields := make([]ticketField, 0)

	for _, rule := range rules {
		r, _ := regexp.Compile(`^([\w\s]+): (\d+)-(\d+) or (\d+)-(\d+)$`)
		digits := r.FindStringSubmatch(rule)
		fieldName := digits[1]
		minOne, _ := strconv.Atoi(digits[2])
		maxOne, _ := strconv.Atoi(digits[3])
		minTwo, _ := strconv.Atoi(digits[4])
		maxTwo, _ := strconv.Atoi(digits[5])
		ticketFields = append(ticketFields, ticketField{name: fieldName, rangeOne: ruleRange{min: minOne, max: maxOne}, rangeTwo: ruleRange{min: minTwo, max: maxTwo}})
	}

	myTicketString := strings.Split(sections[1], "\n")
	fields := strings.Split(myTicketString[1], ",")
	myTicket := make([]int, 0)
	for _, field := range fields {
		intField, _ := strconv.Atoi(field)
		myTicket = append(myTicket, intField)
	}

	nearbyTickets := strings.Split(sections[2], "\n")
	validTickets := make([][]int, 0)
	for i := 1; i < len(nearbyTickets); i++ {
		ticket := nearbyTickets[i]
		fields := strings.Split(ticket, ",")
		intFields := make([]int, 0)
		for _, field := range fields {
			intField, _ := strconv.Atoi(field)
			intFields = append(intFields, intField)
		}

		valid := true
		for _, field := range intFields {
			validField := false
			for _, ticketField := range ticketFields {
				if field <= ticketField.rangeOne.max && field >= ticketField.rangeOne.min {
					validField = true
					break
				}
				if field <= ticketField.rangeTwo.max && field >= ticketField.rangeTwo.min {
					validField = true
					break
				}
			}
			if validField == false {
				valid = false
				break
			}
		}
		if valid == true {
			validTickets = append(validTickets, intFields)
		}
	}

	for tidx, ticketField := range ticketFields {
		for idx := 0; idx < len(ticketFields); idx++ { //for every ticket rule, we need to check every ticket field of the valid tickets
			valid := true
			for _, validTicket := range validTickets {
				if validTicket[idx] <= ticketField.rangeOne.max && validTicket[idx] >= ticketField.rangeOne.min {
					continue
				}
				if validTicket[idx] <= ticketField.rangeTwo.max && validTicket[idx] >= ticketField.rangeTwo.min {
					continue
				}
				valid = false
			}
			if valid == true {
				ticketFields[tidx].positions = append(ticketFields[tidx].positions, idx)
			}
		}
	}

	//fields can be valid for multiple, so we need to eliminate other options to determine their correct location
	tidx := 0
	claimedPositions := make(map[int]bool)
	for tidx < len(ticketFields) {
		ticketField := ticketFields[tidx]
		if len(ticketField.positions) == 1 && claimedPositions[ticketField.positions[0]] == false {
			takenPosition := ticketField.positions[0]
			for oidx, otherTicketField := range ticketFields {
				if oidx != tidx {
					newPositions := make([]int, 0)
					for _, oldPos := range otherTicketField.positions {
						if oldPos != takenPosition {
							newPositions = append(newPositions, oldPos)
						}
					}
					ticketFields[oidx].positions = newPositions
				}
			}
			claimedPositions[takenPosition] = true
			tidx = 0
			continue
		}
		tidx++
	}

	result := 1
	for _, ticketField := range ticketFields {
		if strings.HasPrefix(ticketField.name, "departure") {
			result = result * myTicket[ticketField.positions[0]]
		}
	}
	return result
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
