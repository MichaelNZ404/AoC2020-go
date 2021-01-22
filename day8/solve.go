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

var acc int

type instruction struct {
	phrase   string
	quantity int
}

func solveP1(inputString string) int {
	acc = 0
	r, _ := regexp.Compile(`^(\w{3})\s(.\d+)$`)
	rawInstructions := strings.Split(inputString, "\n")

	// build instruction array
	instructions := make([]instruction, 0)
	for _, rawInstruction := range rawInstructions {
		phrase := r.FindStringSubmatch(rawInstruction)[1]
		quantity, _ := strconv.Atoi(r.FindStringSubmatch(rawInstruction)[2])
		instructions = append(instructions, instruction{phrase: phrase, quantity: quantity})
	}

	//execute
	linesSeen := make(map[int]bool)
	currentLineNum := 0
	for linesSeen[currentLineNum] != true {
		linesSeen[currentLineNum] = true
		if instructions[currentLineNum].phrase == "nop" {
			currentLineNum = currentLineNum + 1
			continue
		}
		if instructions[currentLineNum].phrase == "acc" {
			acc = acc + instructions[currentLineNum].quantity
			currentLineNum = currentLineNum + 1
			continue
		}
		if instructions[currentLineNum].phrase == "jmp" {
			currentLineNum = currentLineNum + instructions[currentLineNum].quantity
			continue
		}
	}
	return acc
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
