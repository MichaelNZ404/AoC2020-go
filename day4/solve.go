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

func checkBYR(passportMap map[string]string) bool {
	if passportMap["byr"] == "" {
		return false
	}
	birthyear, err := strconv.Atoi(passportMap["byr"])
	if err != nil {
		panic(err)
	}
	if birthyear > 2002 || birthyear < 1920 {
		return false
	}
	return true
}

func checkIYR(passportMap map[string]string) bool {
	if passportMap["iyr"] == "" {
		return false
	}
	issueyear, err := strconv.Atoi(passportMap["iyr"])
	if err != nil {
		panic(err)
	}
	if issueyear > 2020 || issueyear < 2010 {
		return false
	}
	return true
}

func checkEYR(passportMap map[string]string) bool {
	if passportMap["eyr"] == "" {
		return false
	}
	expirationyear, err := strconv.Atoi(passportMap["eyr"])
	if err != nil {
		panic(err)
	}
	if expirationyear > 2030 || expirationyear < 2020 {
		return false
	}
	return true
}

func checkHGT(passportMap map[string]string) bool {
	if passportMap["hgt"] == "" {
		return false
	}
	heightNum, err := strconv.Atoi(passportMap["hgt"][:len(passportMap["hgt"])-2])
	if err != nil {
		return false
	}
	heightUnit := passportMap["hgt"][len(passportMap["hgt"])-2:]
	if heightUnit == "cm" {
		if heightNum >= 150 && heightNum <= 193 {
			return true
		}
		return false
	}
	if heightUnit == "in" {
		if heightNum >= 59 && heightNum <= 76 {
			return true
		}
		return false
	}
	return false
}

func checkHCL(passportMap map[string]string) bool {
	if passportMap["hcl"] == "" {
		return false
	}
	match, _ := regexp.MatchString("^#([a-f]|[0-9]){6}$", passportMap["hcl"])
	return match
}

func checkECL(passportMap map[string]string) bool {
	if passportMap["ecl"] == "" {
		return false
	}
	match, _ := regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", passportMap["ecl"])
	return match
}

func checkPID(passportMap map[string]string) bool {
	if passportMap["pid"] == "" {
		return false
	}
	match, _ := regexp.MatchString(`^\d{9}$`, passportMap["pid"])
	return match
}

func isValidPassportStrict(passportString string) bool {
	fields := strings.Fields(passportString)
	passportMap := make(map[string]string)
	for _, field := range fields {
		split := strings.Split(field, ":")
		passportMap[split[0]] = split[1]
	}
	if !checkBYR(passportMap) {
		return false
	}
	if !checkIYR(passportMap) {
		return false
	}
	if !checkEYR(passportMap) {
		return false
	}
	if !checkHGT(passportMap) {
		return false
	}
	if !checkHCL(passportMap) {
		return false
	}
	if !checkECL(passportMap) {
		return false
	}
	if !checkPID(passportMap) {
		return false
	}
	return true
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

func solveP2(inputString string) int {
	passportStrings := cleanPassportData(inputString)
	validCount := 0
	for _, passportString := range passportStrings {
		if isValidPassportStrict(passportString) {
			validCount = validCount + 1
		}
	}
	return validCount
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
