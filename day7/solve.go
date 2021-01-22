package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func readInput(filename string) (returnString string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func sumParents(bagMap map[string]map[string]bool, bagName string) map[string]bool {
	uniqueBags := make(map[string]bool)
	fmt.Println(bagName, bagMap[bagName])
	for childBag := range bagMap[bagName] {
		uniqueBags[childBag] = true
		if bagMap[childBag] == nil {
			continue
		}
		for subchild := range sumParents(bagMap, childBag) {
			uniqueBags[subchild] = true
		}
	}
	return uniqueBags
}

func solveP1(inputString string) int {
	bagMap := make(map[string]map[string]bool)
	rules := strings.Split(inputString, "\n")
	for _, rule := range rules {
		sep := strings.Split(rule, " contain ")
		parentBag := strings.ReplaceAll(sep[0], " bags", "")
		children := strings.Split(sep[1], ", ")
		for _, child := range children {
			if child == "no other bags." {
				continue
			}
			r, _ := regexp.Compile(`^\d\s(\w*\s\w*)\sbags?.?$`)
			childName := r.FindStringSubmatch(child)[1]
			if bagMap[childName] == nil {
				bagMap[childName] = make(map[string]bool)
			}
			bagMap[childName][parentBag] = true
		}
	}
	return len(sumParents(bagMap, "shiny gold"))
}

// func solveP2(inputString string) int {
// 	answerCount := 0
// 	groups := strings.Split(inputString, "\n\n")
// 	for _, group := range groups {
// 		answerMap := make(map[rune]int)
// 		members := strings.Split(group, "\n")
// 		for _, member := range members {
// 			for _, answer := range member {
// 				answerMap[answer] = answerMap[answer] + 1
// 			}
// 		}
// 		for answer := range answerMap {
// 			if answerMap[answer] == len(members) {
// 				answerCount = answerCount + 1
// 			}
// 		}
// 	}
// 	return answerCount
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
