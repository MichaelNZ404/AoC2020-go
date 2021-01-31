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

type treeNode struct {
	a *treeNode
	b *treeNode
}

func buildRuleMap(ruleMapRaw map[string]string, ruleMap map[string][]string, ruleNumber string) {
	ruleString := ruleMapRaw[ruleNumber]
	if ruleString == `"a"` {
		ruleMap[ruleNumber] = []string{"a"}
		return
	}
	if ruleString == `"b"` {
		ruleMap[ruleNumber] = []string{"b"}
		return
	}

	parsedRules := make([]string, 0)
	ruleOptions := strings.Split(ruleString, " | ")
	for _, ruleOption := range ruleOptions {
		subrules := strings.Split(ruleOption, " ")
		ruleOptionResults := make([]string, 0)
		for _, subrule := range subrules {
			if ruleMap[subrule] == nil {
				buildRuleMap(ruleMapRaw, ruleMap, subrule)
			}
			subruleResults := ruleMap[subrule]
			if len(ruleOptionResults) == 0 {
				ruleOptionResults = subruleResults
			} else {
				joinedResults := make([]string, 0)
				for _, firstRule := range ruleOptionResults {
					for _, secondRule := range subruleResults {
						newRule := firstRule + secondRule
						joinedResults = append(joinedResults, newRule)
					}
				}
				ruleOptionResults = joinedResults
			}
		}
		parsedRules = append(parsedRules, ruleOptionResults...)
	}
	ruleMap[ruleNumber] = parsedRules
}

func buildTreeForRule(parent *treeNode, nodes string) {
	nodeValue := nodes[0]
	newLeaf := treeNode{}
	leaf := &newLeaf
	if nodeValue == 'a' {
		if parent.a == nil {
			parent.a = leaf
		} else {
			leaf = parent.a
		}
	}
	if nodeValue == 'b' {
		if parent.b == nil {
			parent.b = leaf
		} else {
			leaf = parent.b
		}
	}
	remainingNodes := nodes[1:]
	if len(remainingNodes) > 0 {
		buildTreeForRule(leaf, remainingNodes)
	}
}

func solveP1(inputString string) int {
	sections := strings.Split(inputString, "\n\n")

	ruleMapRaw := make(map[string]string)
	r, _ := regexp.Compile(`^(\d+): (.+)$`)
	for _, rule := range strings.Split(sections[0], "\n") {
		ruleNumber := r.FindStringSubmatch(rule)[1]
		rule := r.FindStringSubmatch(rule)[2]
		ruleMapRaw[ruleNumber] = rule
	}

	// parse the rules into their literal values
	ruleMap := make(map[string][]string)
	buildRuleMap(ruleMapRaw, ruleMap, "0")

	// build a tree for faster eval
	root := treeNode{}
	for _, rule := range ruleMap["0"] {
		buildTreeForRule(&root, rule)
	}

	validMessageCount := 0
	for _, message := range strings.Split(sections[1], "\n") {
		currentNode := root
		valid := true
		for _, character := range message {
			if character == 'a' && currentNode.a != nil {
				currentNode = *currentNode.a
				continue
			}
			if character == 'b' && currentNode.b != nil {
				currentNode = *currentNode.b
				continue
			}
			valid = false
			break
		}
		if valid {
			validMessageCount++
		}
	}
	return validMessageCount
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
