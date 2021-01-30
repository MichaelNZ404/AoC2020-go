package main

import (
	"fmt"
	"io/ioutil"
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

func solveMath(equation string) int {
	// solve bracket subproblems recursively
	firstBracketIdx := strings.Index(equation, "(")
	for firstBracketIdx >= 0 {
		additionalBracketCount := 0
		lastBracketIdx := -1
		for eidx := firstBracketIdx + 1; eidx < len(equation); eidx++ {
			cString := string(equation[eidx])
			if cString == "(" {
				additionalBracketCount++
				continue
			}
			if cString == ")" && additionalBracketCount > 0 {
				additionalBracketCount--
				continue
			}
			if cString == ")" && additionalBracketCount == 0 {
				lastBracketIdx = eidx
				break
			}
		}
		subProblemResult := solveMath(equation[firstBracketIdx+1 : lastBracketIdx])
		equation = equation[:firstBracketIdx] + fmt.Sprint(subProblemResult) + equation[lastBracketIdx+1:]
		firstBracketIdx = strings.Index(equation, "(")
	}

	//solve bracketless equation
	equationComponents := strings.Split(equation, " ")
	for len(equationComponents) >= 3 {
		v1, _ := strconv.Atoi(equationComponents[0])
		opp := equationComponents[1]
		v2, _ := strconv.Atoi(equationComponents[2])

		switch opp {
		case "+":
			equationComponents = append([]string{fmt.Sprint(v1 + v2)}, equationComponents[3:]...)
			break
		case "*":
			equationComponents = append([]string{fmt.Sprint(v1 * v2)}, equationComponents[3:]...)
			break
		}
	}
	result, _ := strconv.Atoi(equationComponents[0])
	return result
}

func solveAdvancedMath(equation string) int {
	// solve bracket subproblems recursively
	firstBracketIdx := strings.Index(equation, "(")
	for firstBracketIdx >= 0 {
		additionalBracketCount := 0
		lastBracketIdx := -1
		for eidx := firstBracketIdx + 1; eidx < len(equation); eidx++ {
			cString := string(equation[eidx])
			if cString == "(" {
				additionalBracketCount++
				continue
			}
			if cString == ")" && additionalBracketCount > 0 {
				additionalBracketCount--
				continue
			}
			if cString == ")" && additionalBracketCount == 0 {
				lastBracketIdx = eidx
				break
			}
		}
		subProblemResult := solveAdvancedMath(equation[firstBracketIdx+1 : lastBracketIdx])
		equation = equation[:firstBracketIdx] + fmt.Sprint(subProblemResult) + equation[lastBracketIdx+1:]
		firstBracketIdx = strings.Index(equation, "(")
	}

	//solve bracketless equation
	equationComponents := strings.Split(equation, " ")
	for len(equationComponents) >= 3 {
		startIndex := 0
		for i := 0; i < len(equationComponents); i++ {
			if equationComponents[i] == "+" {
				startIndex = i - 1
				break
			}
		}
		v1, _ := strconv.Atoi(equationComponents[startIndex])
		opp := equationComponents[startIndex+1]
		v2, _ := strconv.Atoi(equationComponents[startIndex+2])
		preComponents := equationComponents[:startIndex]
		postComponents := equationComponents[startIndex+3:]
		switch opp {
		case "+":
			equationComponents = append(preComponents, fmt.Sprint(v1+v2))
			break
		case "*":
			equationComponents = append(preComponents, fmt.Sprint(v1*v2))
			break
		}
		equationComponents = append(equationComponents, postComponents...)
	}
	result, _ := strconv.Atoi(equationComponents[0])
	return result
}

func solveP1(inputString string) int {
	mathProblems := strings.Split(inputString, "\n")
	count := 0
	for _, mathProblem := range mathProblems {
		count = count + solveMath(mathProblem)
	}
	return count
}

func solveP2(inputString string) int {
	mathProblems := strings.Split(inputString, "\n")
	count := 0
	for _, mathProblem := range mathProblems {
		count = count + solveAdvancedMath(mathProblem)
	}
	return count
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
