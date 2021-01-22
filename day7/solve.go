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

func sumParents(bagMap map[string]map[string]bool, bagName string) map[string]bool {
	uniqueBags := make(map[string]bool)
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

// In part one we build a map where the child bag is the key and the parent bags are the children.
// We do this so we can easily see which bag is contained by other bags, rather than searching the entire map
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

type bagCount struct {
	count int
	name  string
}

func countChildren(bagMap map[string][]bagCount, bagName string) int {
	count := 0
	for _, childbag := range bagMap[bagName] {
		if bagMap[childbag.name] == nil {
			count = count + childbag.count
		} else {
			count = count + (childbag.count * countChildren(bagMap, childbag.name))
		}
	}
	return count + 1
}

// in part two the map is in the more sensible order (parent bag is the key) - however here we must use a struct for the children to contain the count.
func solveP2(inputString string) int {
	bagMap := make(map[string][]bagCount)
	rules := strings.Split(inputString, "\n")
	for _, rule := range rules {
		sep := strings.Split(rule, " contain ")
		parentBag := strings.ReplaceAll(sep[0], " bags", "")
		children := strings.Split(sep[1], ", ")
		for _, child := range children {
			if child == "no other bags." {
				continue
			}
			r, _ := regexp.Compile(`^(\d)\s(\w*\s\w*)\sbags?.?$`)
			childCount, _ := strconv.Atoi(r.FindStringSubmatch(child)[1])
			childName := r.FindStringSubmatch(child)[2]

			if bagMap[parentBag] == nil {
				bagMap[parentBag] = make([]bagCount, 0)
			}
			bagMap[parentBag] = append(bagMap[parentBag], bagCount{count: childCount, name: childName})
		}
	}

	return countChildren(bagMap, "shiny gold") - 1 //minus one for self
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
