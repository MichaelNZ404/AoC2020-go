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

func countSeatIfOccupied(seat byte) int {
	if seat == '#' {
		return 1
	}
	return 0
}

func getAdjacentOccupiedSeatCount(rowIdx int, seatIdx int, lastState []string) int {
	adjacentCount := 0
	if rowIdx > 0 {
		if seatIdx > 0 {
			adjacentCount = adjacentCount + countSeatIfOccupied(lastState[rowIdx-1][seatIdx-1]) //up-left
		}
		adjacentCount = adjacentCount + countSeatIfOccupied(lastState[rowIdx-1][seatIdx]) //up
		if seatIdx+1 < len(lastState[rowIdx]) {
			adjacentCount = adjacentCount + countSeatIfOccupied(lastState[rowIdx-1][seatIdx+1]) //up-right
		}
	}

	if seatIdx > 0 {
		adjacentCount = adjacentCount + countSeatIfOccupied(lastState[rowIdx][seatIdx-1]) //left
	}
	if seatIdx+1 < len(lastState[rowIdx]) {
		adjacentCount = adjacentCount + countSeatIfOccupied(lastState[rowIdx][seatIdx+1]) //right
	}

	if rowIdx+1 < len(lastState) {
		if seatIdx > 0 {
			adjacentCount = adjacentCount + countSeatIfOccupied(lastState[rowIdx+1][seatIdx-1]) //down-left
		}
		adjacentCount = adjacentCount + countSeatIfOccupied(lastState[rowIdx+1][seatIdx]) //down
		if seatIdx+1 < len(lastState[rowIdx+1]) {
			adjacentCount = adjacentCount + countSeatIfOccupied(lastState[rowIdx+1][seatIdx+1]) //down-right
		}
	}
	return adjacentCount
}

func solveP1(inputString string) int {
	seatRows := strings.Split(inputString, "\n")
	nextState := seatRows
	lastState := make([]string, len(nextState))
	occupiedSeatCount := 0
	for fmt.Sprint(lastState) != fmt.Sprint(nextState) {
		// reset last state
		for idx, row := range nextState {
			lastState[idx] = row
		}

		// run the update
		for rowIdx, row := range lastState {
			for seatIdx, seat := range row {
				if seat == '.' {
					continue
				}
				if seat == 'L' {
					if getAdjacentOccupiedSeatCount(rowIdx, seatIdx, lastState) == 0 {
						occupiedSeatCount++
						nextState[rowIdx] = nextState[rowIdx][:seatIdx] + "#" + nextState[rowIdx][seatIdx+1:]
					}
					continue
				}
				if seat == '#' {
					if getAdjacentOccupiedSeatCount(rowIdx, seatIdx, lastState) >= 4 {
						occupiedSeatCount--
						nextState[rowIdx] = nextState[rowIdx][:seatIdx] + "L" + nextState[rowIdx][seatIdx+1:]
					}
					continue
				}
			}
		}
	}
	return occupiedSeatCount
}

func getAdjacentOccupiedSeatCountLong(rowIdx int, seatIdx int, lastState []string) int {
	adjacentCount := 0
	for idx, jdx := rowIdx-1, seatIdx-1; idx >= 0 && jdx >= 0; idx, jdx = idx-1, jdx-1 { //up-left
		if lastState[idx][jdx] == '#' {
			adjacentCount++
			break
		}
		if lastState[idx][jdx] == 'L' {
			break
		}
	}
	for idx := rowIdx - 1; idx >= 0; idx-- { //up
		if lastState[idx][seatIdx] == '#' {
			adjacentCount++
			break
		}
		if lastState[idx][seatIdx] == 'L' {
			break
		}
	}
	for idx, jdx := rowIdx-1, seatIdx+1; idx >= 0 && jdx < len(lastState[rowIdx]); idx, jdx = idx-1, jdx+1 { //up-right
		if lastState[idx][jdx] == '#' {
			adjacentCount++
			break
		}
		if lastState[idx][jdx] == 'L' {
			break
		}
	}
	for jdx := seatIdx - 1; jdx >= 0; jdx-- { //left
		if lastState[rowIdx][jdx] == '#' {
			adjacentCount++
			break
		}
		if lastState[rowIdx][jdx] == 'L' {
			break
		}
	}
	for jdx := seatIdx + 1; jdx < len(lastState[rowIdx]); jdx++ { //right
		if lastState[rowIdx][jdx] == '#' {
			adjacentCount++
			break
		}
		if lastState[rowIdx][jdx] == 'L' {
			break
		}
	}
	for idx, jdx := rowIdx+1, seatIdx-1; idx < len(lastState) && jdx >= 0; idx, jdx = idx+1, jdx-1 { //down-left
		if lastState[idx][jdx] == '#' {
			adjacentCount++
			break
		}
		if lastState[idx][jdx] == 'L' {
			break
		}
	}
	for idx := rowIdx + 1; idx < len(lastState); idx++ { //down
		if lastState[idx][seatIdx] == '#' {
			adjacentCount++
			break
		}
		if lastState[idx][seatIdx] == 'L' {
			break
		}
	}
	for idx, jdx := rowIdx+1, seatIdx+1; idx < len(lastState) && jdx < len(lastState[rowIdx]); idx, jdx = idx+1, jdx+1 { //down-right
		if lastState[idx][jdx] == '#' {
			adjacentCount++
			break
		}
		if lastState[idx][jdx] == 'L' {
			break
		}
	}
	return adjacentCount
}

func solveP2(inputString string) int {
	seatRows := strings.Split(inputString, "\n")
	nextState := seatRows
	lastState := make([]string, len(nextState))
	occupiedSeatCount := 0
	for fmt.Sprint(lastState) != fmt.Sprint(nextState) {
		// reset last state
		for idx, row := range nextState {
			lastState[idx] = row
		}

		// run the update
		for rowIdx, row := range lastState {
			for seatIdx, seat := range row {
				if seat == '.' {
					continue
				}
				if seat == 'L' {
					if getAdjacentOccupiedSeatCountLong(rowIdx, seatIdx, lastState) == 0 {
						occupiedSeatCount++
						nextState[rowIdx] = nextState[rowIdx][:seatIdx] + "#" + nextState[rowIdx][seatIdx+1:]
					}
					continue
				}
				if seat == '#' {
					if getAdjacentOccupiedSeatCountLong(rowIdx, seatIdx, lastState) >= 5 {
						occupiedSeatCount--
						nextState[rowIdx] = nextState[rowIdx][:seatIdx] + "L" + nextState[rowIdx][seatIdx+1:]
					}
					continue
				}
			}
		}
	}
	return occupiedSeatCount
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
