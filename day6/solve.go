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

// func solveP2(inputString string) int {
// 	highestSeat := 128 * 8
// 	seatsTaken := make([]bool, highestSeat)
// 	passes := strings.Split(inputString, "\n")
// 	for _, pass := range passes {
// 		candidateRows := makeRange(0, 127)
// 		candidateColumns := makeRange(0, 7)
// 		for _, letter := range pass {
// 			switch letter {
// 			case 'F':
// 				candidateRows = candidateRows[:len(candidateRows)/2]
// 			case 'B':
// 				candidateRows = candidateRows[len(candidateRows)/2:]
// 			case 'L':
// 				candidateColumns = candidateColumns[:len(candidateColumns)/2]
// 			case 'R':
// 				candidateColumns = candidateColumns[len(candidateColumns)/2:]
// 			}
// 		}
// 		if len(candidateRows) != 1 || len(candidateColumns) != 1 {
// 			panic(errors.New("Filtering did not narrow to a single seat"))
// 		}
// 		seatNum := (candidateRows[0] * 8) + candidateColumns[0]
// 		seatsTaken[seatNum] = true
// 	}
// 	for idx, isTaken := range seatsTaken {
// 		if idx > 0 && idx < highestSeat {
// 			if !isTaken && seatsTaken[idx-1] && seatsTaken[idx+1] {
// 				return idx
// 			}
// 		}
// 	}
// 	panic(errors.New("Unable to find empty seat"))
// }

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
