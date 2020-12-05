package main

import (
	"fmt"
	"aoc-2020/utils"
)

const tree = "#"

type slope struct {
	right int
	down  int
}

func main() {
	text := utils.ReadTextFile("../day_03/data.txt")

	slopeArrSolution1 := []slope{{3, 1}}
	slopeArrSolution2 := []slope{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	fmt.Println("solution 1: ", getMultipliedTreesFromSlopes(slopeArrSolution1, text))
	fmt.Println("solution 2: ", getMultipliedTreesFromSlopes(slopeArrSolution2, text))
}

func getMultipliedTreesFromSlopes(slopes []slope, inputMap []string) int {
	numOfTrees := 1

	for _, slopeItem := range slopes {
		numOfTrees = numOfTrees * getTreesFromSlope(slopeItem, inputMap)
	}

	return numOfTrees
}

func getTreesFromSlope(slopeItem slope, inputMap []string) int {
	rowLength := len(inputMap[0])
	numOfTrees := 0
	x := 0
	y := 0

	for i := 0; i < (len(inputMap) / slopeItem.down - 1); i++ {
		x = (x + slopeItem.right) % rowLength
		y = y + slopeItem.down

		if string(inputMap[y][x]) == tree {
			numOfTrees = numOfTrees + 1
		}
	}

	return numOfTrees
}