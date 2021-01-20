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

func cleanPassportData(inputString string) []string {
	passportStrings := strings.Split(inputString, "\n\n")
	cleanPassportStrings := make([]string, 0)
	for _, passportString := range passportStrings {
		cleanPassportStrings = append(cleanPassportStrings, strings.ReplaceAll(passportString, "\n", " "))
	}
	return cleanPassportStrings
}

func isValidPassport(passportString string) bool {
	// fmt.Println(passportString)
	fields := strings.Fields(passportString)
	passportMap := make(map[string]string)
	for _, field := range fields {
		split := strings.Split(field, ":")
		passportMap[split[0]] = split[1]
	}
	if passportMap["byr"] != "" && passportMap["iyr"] != "" && passportMap["eyr"] != "" && passportMap["hgt"] != "" && passportMap["hcl"] != "" && passportMap["ecl"] != "" && passportMap["pid"] != "" {
		return true
	}
	return false
}

func solveP1(inputString string) int {
	passportStrings := cleanPassportData(inputString)
	validCount := 0
	for _, passportString := range passportStrings {
		if isValidPassport(passportString) {
			validCount = validCount + 1
		}
	}
	return validCount
}

// func solveP2(stringList []string) int {
// 	totalHit := 1
// 	totalHit = totalHit * checkSlope(stringList, 1, 1)
// 	totalHit = totalHit * checkSlope(stringList, 3, 1)
// 	totalHit = totalHit * checkSlope(stringList, 5, 1)
// 	totalHit = totalHit * checkSlope(stringList, 7, 1)
// 	totalHit = totalHit * checkSlope(stringList, 1, 2)
// 	return totalHit
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
