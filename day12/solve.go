package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

type instructionStruct struct {
	command string
	value   int
}

func solveP1(inputString string) int {
	lines := strings.Split(inputString, "\n")
	r, _ := regexp.Compile(`^(\w)(\d+)$`)
	instructions := make([]instructionStruct, len(lines))
	for idx, line := range lines {
		command := r.FindStringSubmatch(line)[1]
		value, _ := strconv.Atoi(r.FindStringSubmatch(line)[2])
		instructions[idx] = instructionStruct{command: command, value: value}
	}

	shipDir := 90 //assuming roations are multiples of 90 degrees
	shipX, shipY := 0, 0

	for _, instruction := range instructions {
		switch instruction.command {
		case "F":
			switch shipDir {
			case 0:
				shipY += instruction.value
				break
			case 90:
				shipX += instruction.value
				break
			case 180:
				shipY -= instruction.value
				break
			case 270:
				shipX -= instruction.value
				break
			}
			break
		case "N":
			shipY += instruction.value
			break
		case "S":
			shipY -= instruction.value
			break
		case "E":
			shipX += instruction.value
			break
		case "W":
			shipX -= instruction.value
			break
		case "L":
			shipDir = (shipDir - instruction.value) % 360
			if shipDir < 0 {
				shipDir += 360
			}
			break
		case "R":
			shipDir = (shipDir + instruction.value) % 360
			if shipDir < 0 {
				shipDir += 360
			}
			break
		}
	}
	return int(math.Abs(float64(shipX)) + math.Abs(float64(shipY)))
}

func rotateRight(x int, y int, quantity int) (returnX int, returnY int) {
	for i := quantity / 90; i > 0; i-- {
		oldX := x
		oldY := y
		x = oldY
		y = -oldX
	}
	return x, y
}

func rotateLeft(x int, y int, quantity int) (returnX int, returnY int) {
	for i := quantity / 90; i > 0; i-- {
		oldX := x
		oldY := y
		x = -oldY
		y = oldX
	}
	return x, y
}

func solveP2(inputString string) int {
	lines := strings.Split(inputString, "\n")
	r, _ := regexp.Compile(`^(\w)(\d+)$`)
	instructions := make([]instructionStruct, len(lines))
	for idx, line := range lines {
		command := r.FindStringSubmatch(line)[1]
		value, _ := strconv.Atoi(r.FindStringSubmatch(line)[2])
		instructions[idx] = instructionStruct{command: command, value: value}
	}

	shipX, shipY := 0, 0
	waypointOffsetX, waypointOffsetY := 10, 1

	for _, instruction := range instructions {
		switch instruction.command {
		case "F":
			shipX += (waypointOffsetX * instruction.value)
			shipY += (waypointOffsetY * instruction.value)
			break
		case "N":
			waypointOffsetY += instruction.value
			break
		case "S":
			waypointOffsetY -= instruction.value
			break
		case "E":
			waypointOffsetX += instruction.value
			break
		case "W":
			waypointOffsetX -= instruction.value
			break
		case "L":
			waypointOffsetX, waypointOffsetY = rotateLeft(waypointOffsetX, waypointOffsetY, instruction.value)
			break
		case "R":
			waypointOffsetX, waypointOffsetY = rotateRight(waypointOffsetX, waypointOffsetY, instruction.value)
			break
		}
	}
	return int(math.Abs(float64(shipX)) + math.Abs(float64(shipY)))
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
