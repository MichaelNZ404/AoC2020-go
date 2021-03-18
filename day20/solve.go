package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

type cameraTile struct {
	id    int
	image [][]rune
}

func (x cameraTile) rotate() {
	newImage := make([][]rune, len(x.image))
	for i := range newImage {
		newImage[i] = make([]rune, len(x.image))
	}

	for ridx, row := range x.image {
		for cidx, col := range row {
			newImage[cidx][len(x.image)-ridx-1] = col
		}
	}
	for i := range newImage {
		x.image[i] = newImage[i]
	}
}

func (x cameraTile) flip() {
	newImage := make([][]rune, len(x.image))
	for ridx, row := range x.image {
		newImage[len(x.image)-ridx-1] = row
	}
	for i := range x.image {
		x.image[i] = newImage[i]
	}
}
func (x cameraTile) getTop() []rune {
	return x.image[0]
}
func (x cameraTile) getBottom() []rune {
	return x.image[len(x.image)-1]
}
func (x cameraTile) getRight() []rune {
	side := make([]rune, len(x.image))
	for i := 0; i < len(x.image); i++ {
		side[i] = x.image[i][len(x.image)-1]
	}
	return side
}
func (x cameraTile) getLeft() []rune {
	side := make([]rune, len(x.image))
	for i := 0; i < len(x.image); i++ {
		side[i] = x.image[i][0]
	}
	return side
}

func reverseString(in string) string {
	out := ""
	for _, r := range in {
		out = string(r) + out
	}
	return out
}

// func solveP1(inputString string) int {
// 	tiles := strings.Split(inputString, "\n\n")
// 	sideMap := make(map[string][]cameraTile)

// 	for _, tile := range tiles {
// 		rows := strings.Split(tile, "\n")
// 		idRegex, _ := regexp.Compile(`Tile (\d+)`)
// 		tileID, _ := strconv.Atoi(idRegex.FindStringSubmatch(rows[0])[1])

// 		top := rows[1]
// 		bottom := rows[len(rows)-1]
// 		left := ""
// 		right := ""
// 		for i := 1; i < len(rows); i++ {
// 			left = left + string(rows[i][0])
// 			right = right + string(rows[i][len(top)-1])
// 		}

// 		unflippedTile := cameraTile{id: tileID, sides: [4]string{top, right, left, bottom}}
// 		flippedTile := cameraTile{id: tileID, sides: [4]string{reverseString(top), reverseString(right), reverseString(left), reverseString(bottom)}, flipped: true}

// 		for _, side := range unflippedTile.sides {
// 			sideMap[side] = append(sideMap[side], unflippedTile)
// 		}
// 		for _, side := range flippedTile.sides {
// 			sideMap[side] = append(sideMap[side], flippedTile)
// 		}
// 	}

// 	for key, value := range sideMap {
// 		fmt.Println(key, len(value))
// 	}
// 	return 1
// }

func toIds(tiles []cameraTile) []int {
	returnIds := make([]int, 0)
	for _, val := range tiles {
		returnIds = append(returnIds, val.id)
	}
	return returnIds
}

func solve(pickedTiles []cameraTile, unpickedTiles []cameraTile, imageLength int) ([]cameraTile, []cameraTile) {
	currentIdxToSolve := len(pickedTiles)
	// fmt.Println("Searching solution with ", toIds(pickedTiles), toIds(unpickedTiles), currentIdxToSolve)

	for idx, unplacedTile := range unpickedTiles {
		for i := 0; i < 2; i++ {
			for j := 0; j < 4; j++ {
				if (currentIdxToSolve % imageLength) != 0 { // not in first column
					if string(pickedTiles[len(pickedTiles)-1].getRight()) != string(unplacedTile.getLeft()) {
						unplacedTile.rotate()
						continue
					}
				}
				if currentIdxToSolve >= imageLength { // not in first row
					if string(pickedTiles[currentIdxToSolve-imageLength].getBottom()) != string(unplacedTile.getTop()) {
						unplacedTile.rotate()
						continue
					}
				}
				// attempt solve
				potentialSolution := append(pickedTiles, unplacedTile)
				remainingTiles := make([]cameraTile, 0)
				for i, t := range unpickedTiles {
					if i == idx {
						continue
					}
					remainingTiles = append(remainingTiles, t)
				}
				// fmt.Println("Removing ", unplacedTile.id, "...", toIds(unpickedTiles), toIds(remainingTiles))
				upstreamSolution, upstreamRemainder := solve(potentialSolution, remainingTiles, imageLength)
				if len(upstreamRemainder) == 0 {
					return upstreamSolution, upstreamRemainder
				}
				//end solve attempt
			}
			unplacedTile.flip()
		}
	}
	return pickedTiles, unpickedTiles
}

func solveP1(inputString string) int {
	unparsedTiles := strings.Split(inputString, "\n\n")
	imageLength := int(math.Sqrt(float64(len(unparsedTiles))))
	unpickedTileList := make([]cameraTile, 0)
	pickedTileList := make([]cameraTile, 0)

	for _, unparsedTile := range unparsedTiles {
		rows := strings.Split(unparsedTile, "\n")
		idRegex, _ := regexp.Compile(`Tile (\d+)`)
		tileID, _ := strconv.Atoi(idRegex.FindStringSubmatch(rows[0])[1])

		image := make([][]rune, 0)
		for i := 1; i < len(rows); i++ {
			image = append(image, []rune(rows[i]))
		}
		unpickedTileList = append(unpickedTileList, cameraTile{id: tileID, image: image})
	}
	solutionImage, _ := solve(pickedTileList, unpickedTileList, imageLength)
	solution := solutionImage[0].id * solutionImage[imageLength-1].id * solutionImage[len(unparsedTiles)-imageLength].id * solutionImage[len(unparsedTiles)-1].id

	for i := 0; i < imageLength; i++ {
		for rowidx := range solutionImage[0].image {
			fmt.Println(
				string(solutionImage[0+(i*imageLength)].image[rowidx]),
				string(solutionImage[1+(i*imageLength)].image[rowidx]),
				string(solutionImage[2+(i*imageLength)].image[rowidx]),
			)
		}
		fmt.Println(" ")
	}
	return solution
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
