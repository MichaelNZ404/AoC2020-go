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

func solveP1(inputString string) int {
	lines := strings.Split(inputString, "\n")
	var mask string
	maskRegex, _ := regexp.Compile(`mask = (.{36})`)
	memRegex, _ := regexp.Compile(`mem\[(\d+)\] = (\d+)`)
	memory := make(map[int]int)

	for _, line := range lines {
		if maskRegex.MatchString(line) {
			mask = maskRegex.FindStringSubmatch(line)[1]
		}
		if memRegex.MatchString(line) {
			vals := memRegex.FindStringSubmatch(line)
			memoryAddress, _ := strconv.Atoi(vals[1])
			memoryValue, _ := strconv.Atoi(vals[2])
			memoryValueBits := fmt.Sprintf("%b", memoryValue)
			memoryValueBits = strings.Repeat("0", 36-len(memoryValueBits)) + memoryValueBits
			bitArray := strings.Split(memoryValueBits, "")
			for idx, value := range mask {
				if value == '1' || value == '0' {
					bitArray[idx] = string(value)
				}
			}
			maskedValue, _ := strconv.ParseInt(strings.Join(bitArray, ""), 2, 37)
			memory[memoryAddress] = int(maskedValue)
		}
	}

	total := 0
	for _, i := range memory {
		total += i
	}
	return total
}

func resolveFloating(input []string, startLoc int) [][]string {
	results := make([][]string, 0)
	valid := true
	for idx := startLoc; idx < len(input); idx++ {
		value := input[idx]
		if value == "X" {
			valid = false
			inputBranchOne, inputBranchTwo := append([]string(nil), input...), append([]string(nil), input...)
			inputBranchOne[idx], inputBranchTwo[idx] = "1", "0"
			for _, result := range resolveFloating(inputBranchOne, idx+1) {
				results = append(results, result)
			}
			for _, result := range resolveFloating(inputBranchTwo, idx+1) {
				results = append(results, result)
			}
			break
		}
	}
	if valid == true {
		results = append(results, input)
	}
	return results
}

func solveP2(inputString string) int {
	lines := strings.Split(inputString, "\n")
	var mask string
	maskRegex, _ := regexp.Compile(`mask = (.{36})`)
	memRegex, _ := regexp.Compile(`mem\[(\d+)\] = (\d+)`)
	memory := make(map[int]int)

	for _, line := range lines {
		if maskRegex.MatchString(line) {
			mask = maskRegex.FindStringSubmatch(line)[1]
		}
		if memRegex.MatchString(line) {
			vals := memRegex.FindStringSubmatch(line)
			memoryAddress, _ := strconv.Atoi(vals[1])
			memoryValue, _ := strconv.Atoi(vals[2])

			memoryAddressBits := fmt.Sprintf("%b", memoryAddress)
			memoryAddressBits = strings.Repeat("0", 36-len(memoryAddressBits)) + memoryAddressBits
			bitArray := strings.Split(memoryAddressBits, "")
			for idx, value := range mask {
				if value == '1' || value == 'X' {
					bitArray[idx] = string(value)
				}
			}

			addresses := resolveFloating(bitArray, 0)
			for _, addr := range addresses {
				castedValue, _ := strconv.ParseInt(strings.Join(addr, ""), 2, 37)
				memory[int(castedValue)] = memoryValue
			}
		}
	}

	total := 0
	for _, i := range memory {
		total += i
	}
	return total
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
