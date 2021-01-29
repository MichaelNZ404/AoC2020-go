package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
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
	earliestTime, _ := strconv.Atoi(lines[0])

	minTime := 0
	minTimeBusID := 0
	for _, val := range strings.Split(lines[1], ",") {
		if val != "x" {
			busID, _ := strconv.Atoi(val)
			candidateTime := int(math.Ceil(float64(earliestTime)/float64(busID))) * busID
			if minTimeBusID == 0 || candidateTime < minTime {
				minTime = candidateTime
				minTimeBusID = busID
			}
		}
	}
	return minTimeBusID * (minTime - earliestTime)
}

// slow solution
// func solveP2(inputString string) int {
// 	lines := strings.Split(inputString, "\n")

// 	busIds := make([]int, 0)
// 	for _, val := range strings.Split(lines[1], ",") {
// 		if val != "x" {
// 			busID, _ := strconv.Atoi(val)
// 			busIds = append(busIds, busID)
// 		} else {
// 			busIds = append(busIds, 0)
// 		}
// 	}

// 	candidateTime := 0
// 	lastCheckedMult := 0 //this is the last multiple of index 0 checked, keeps our place in the loop
// 	i := 0
// 	for i < len(busIds) {
// 		if i == 0 {
// 			lastCheckedMult++
// 			candidateTime = lastCheckedMult * busIds[0]
// 		} else {
// 			if busIds[i] == 0 { // this is an 'x'
// 				i++
// 				continue
// 			}
// 			if (candidateTime+i)%busIds[i] != 0 { // invalid
// 				i = 0
// 				continue
// 			}
// 		}
// 		i++
// 	}
// 	return candidateTime
// }

//https://rosettacode.org/wiki/Chinese_remainder_theorem#Go
var one = big.NewInt(1)

func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}

func solveP2(inputString string) int {
	lines := strings.Split(inputString, "\n")

	busIds := make([]*big.Int, 0)
	busRemainders := make([]*big.Int, 0)
	for idx, val := range strings.Split(lines[1], ",") {
		if val != "x" {
			busID, _ := strconv.Atoi(val)
			busIds = append(busIds, big.NewInt(int64(busID)))
			busRemainders = append(busRemainders, big.NewInt(int64(-idx)))
		}
	}
	result, _ := crt(busRemainders, busIds)
	return int(result.Int64())
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
