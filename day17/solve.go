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

type threeDimensionSpace struct {
	z int
	y int
	x int
}

type hyperSpace struct {
	z int
	y int
	x int
	w int
}

// count the active neighbours for a current cube. Also, store adjacent inactive cubes in a map so we can check their adjacency conditions later
func getActiveNeighbourCount(cubeMap map[threeDimensionSpace]bool, currentPosition threeDimensionSpace, inactiveMap map[threeDimensionSpace]bool) int {
	count := 0
	for z := currentPosition.z - 1; z <= currentPosition.z+1; z++ {
		for y := currentPosition.y - 1; y <= currentPosition.y+1; y++ {
			for x := currentPosition.x - 1; x <= currentPosition.x+1; x++ {
				if currentPosition.x == x && currentPosition.y == y && currentPosition.z == z {
					continue
				}
				if cubeMap[threeDimensionSpace{x: x, y: y, z: z}] == true {
					count++
				} else {
					inactiveMap[threeDimensionSpace{x: x, y: y, z: z}] = true
				}
			}
		}
	}
	return count
}

// same as getActiveNeighbourCount, but for 4D space
func getHyperActiveNeighbourCount(cubeMap map[hyperSpace]bool, currentPosition hyperSpace, inactiveMap map[hyperSpace]bool) int {
	count := 0
	for w := currentPosition.w - 1; w <= currentPosition.w+1; w++ {
		for z := currentPosition.z - 1; z <= currentPosition.z+1; z++ {
			for y := currentPosition.y - 1; y <= currentPosition.y+1; y++ {
				for x := currentPosition.x - 1; x <= currentPosition.x+1; x++ {
					if currentPosition.x == x && currentPosition.y == y && currentPosition.z == z && currentPosition.w == w {
						continue
					}
					if cubeMap[hyperSpace{x: x, y: y, z: z, w: w}] == true {
						count++
					} else {
						inactiveMap[hyperSpace{x: x, y: y, z: z, w: w}] = true
					}
				}
			}
		}
	}
	return count
}

func solveP1(inputString string) int {
	rows := strings.Split(inputString, "\n")
	activeCubeMap := make(map[threeDimensionSpace]bool)

	z := 0
	for y, row := range rows {
		for x, state := range row {
			if state == '#' { //active
				activeCubeMap[threeDimensionSpace{x: x, y: y, z: z}] = true
			}
		}
	}

	for cycles := 0; cycles < 6; cycles++ {
		updatedCubeMap := make(map[threeDimensionSpace]bool)
		inactiveMap := make(map[threeDimensionSpace]bool)
		inactiveAdjacents := make(map[threeDimensionSpace]bool)
		for location := range activeCubeMap {
			adjacentCount := getActiveNeighbourCount(activeCubeMap, location, inactiveMap)
			if adjacentCount == 2 || adjacentCount == 3 {
				updatedCubeMap[location] = true
			}
		}

		for location := range inactiveMap {
			adjacentCount := getActiveNeighbourCount(activeCubeMap, location, inactiveAdjacents)
			if adjacentCount == 3 {
				updatedCubeMap[location] = true
			}
		}
		activeCubeMap = updatedCubeMap
	}
	return len(activeCubeMap)
}

func solveP2(inputString string) int {
	rows := strings.Split(inputString, "\n")
	activeCubeMap := make(map[hyperSpace]bool)

	z := 0
	w := 0
	for y, row := range rows {
		for x, state := range row {
			if state == '#' { //active
				activeCubeMap[hyperSpace{x: x, y: y, z: z, w: w}] = true
			}
		}
	}

	for cycles := 0; cycles < 6; cycles++ {
		updatedCubeMap := make(map[hyperSpace]bool)
		inactiveMap := make(map[hyperSpace]bool)
		inactiveAdjacents := make(map[hyperSpace]bool)
		for location := range activeCubeMap {
			adjacentCount := getHyperActiveNeighbourCount(activeCubeMap, location, inactiveMap)
			if adjacentCount == 2 || adjacentCount == 3 {
				updatedCubeMap[location] = true
			}
		}

		for location := range inactiveMap {
			adjacentCount := getHyperActiveNeighbourCount(activeCubeMap, location, inactiveAdjacents)
			if adjacentCount == 3 {
				updatedCubeMap[location] = true
			}
		}
		activeCubeMap = updatedCubeMap
	}
	return len(activeCubeMap)
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
