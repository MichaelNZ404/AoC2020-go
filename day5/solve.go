package main

import (
	"errors"
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

// makes a array of sequential numbers
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func solveP1(inputString string) int {
	highestNum := 0
	passes := strings.Split(inputString, "\n")
	for _, pass := range passes {
		candidateRows := makeRange(0, 127)
		candidateColumns := makeRange(0, 7)
		for _, letter := range pass {
			switch letter {
			case 'F':
				candidateRows = candidateRows[:len(candidateRows)/2]
			case 'B':
				candidateRows = candidateRows[len(candidateRows)/2:]
			case 'L':
				candidateColumns = candidateColumns[:len(candidateColumns)/2]
			case 'R':
				candidateColumns = candidateColumns[len(candidateColumns)/2:]
			}
		}
		if len(candidateRows) != 1 || len(candidateColumns) != 1 {
			panic(errors.New("Filtering did not narrow to a single seat"))
		}
		seatNum := (candidateRows[0] * 8) + candidateColumns[0]
		if seatNum > highestNum {
			highestNum = seatNum
		}
	}
	return highestNum
}

func solveP2(inputString string) int {
	highestSeat := 128 * 8
	seatsTaken := make([]bool, highestSeat)
	passes := strings.Split(inputString, "\n")
	for _, pass := range passes {
		candidateRows := makeRange(0, 127)
		candidateColumns := makeRange(0, 7)
		for _, letter := range pass {
			switch letter {
			case 'F':
				candidateRows = candidateRows[:len(candidateRows)/2]
			case 'B':
				candidateRows = candidateRows[len(candidateRows)/2:]
			case 'L':
				candidateColumns = candidateColumns[:len(candidateColumns)/2]
			case 'R':
				candidateColumns = candidateColumns[len(candidateColumns)/2:]
			}
		}
		if len(candidateRows) != 1 || len(candidateColumns) != 1 {
			panic(errors.New("Filtering did not narrow to a single seat"))
		}
		seatNum := (candidateRows[0] * 8) + candidateColumns[0]
		seatsTaken[seatNum] = true
	}
	for idx, isTaken := range seatsTaken {
		if idx > 0 && idx < highestSeat {
			if !isTaken && seatsTaken[idx-1] && seatsTaken[idx+1] {
				return idx
			}
		}
	}
	panic(errors.New("Unable to find empty seat"))
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
