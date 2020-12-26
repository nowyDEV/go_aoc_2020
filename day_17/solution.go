package main

import (
	"aoc-2020/utils"
	"fmt"
	"strings"
)

type cube struct {
	x      int
	y      int
	z      int
	active bool
}

func main() {
	data := utils.ReadTextFile("./data_test.txt")

	cubeMap := generateCubeMap(data)

	fmt.Println(cubeMap)
}

func getNumOfActiveNeighbors(cubeMap []cube, target cube) int {
	result := 0
	for _, item := range cubeMap {
		if item.x == target.x && item.y == target.y && item.z == target.z {
			result++
		}
	}

	return result
}

func generateCubeMap(data []string) []cube {
	var result []cube

	for index, item := range data {
		row := strings.Split(item, "")
		for rowIndex, rowItem := range row {
			isActive := rowItem == "#"
			result = append(result, cube{rowIndex, index, 0, isActive})
		}
	}

	return result
}
