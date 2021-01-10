package main

import (
	"aoc-2020/utils"
	"fmt"
	"sort"
)

func main() {
	data := utils.ReadTextFile("./data.txt")

	instructions := parseData(data)

	var coordinates [][]int

	for i := 0; i < len(data); i++ {
		coordinates = append(coordinates, getCoordinates(instructions[i]))
	}

	sortCoords(coordinates)
	dataList := buildFlipDataList(coordinates)

	fmt.Println("part 1: ", len(dataList) - utils.Sum(dataList))
}

type coords [][]int

func (list coords) Len() int { return len(list) }

func (list coords) Swap(i, j int) { list[i], list[j] = list[j], list[i] }

func (list coords) Less(i, j int) bool {
	var si []int = list[i]
	var sj []int = list[j]
	if si[0] == sj[0] {
		return si[1] < sj[1]
	}
	return si[0] < sj[0]
}

func sortCoords(coordinates [][]int) [][]int {
	sort.Sort(coords(coordinates))

	return coordinates
}

// 0 - black, 1 - white
func buildFlipDataList(sortedCoords [][]int) (result []int) {
	startValue := 0

	for i := 1; i < len(sortedCoords); {
		if areSameCoords(sortedCoords[i-1], sortedCoords[i]) {
			startValue = (startValue - 1) * -1
			i += 1

			if i == len(sortedCoords) {
				result = append(result, startValue)
			}
		} else {
			result = append(result, startValue)
			i += 1
			startValue = 0

			if i == len(sortedCoords) {
				result = append(result, startValue)
			}
		}
	}

	return result
}

func areSameCoords(a, b []int) bool {
	return a[0] == b[0] && a[1] == b[1]
}

func parseData(input []string) (result [][]string) {
	for _, row := range input {
		result = append(result, parseInstruction(row))
	}

	return result
}

func parseInstruction(input string) (result []string) {
	var directions = []string{"e", "se", "sw", "w", "nw", "ne"}

	for i := 0; i < len(input)-1; {
		nextTwo := input[i : i+2]

		if utils.ContainsString(directions, nextTwo) {
			result = append(result, nextTwo)
			i += 2
		} else {
			result = append(result, nextTwo[0:1])
			i += 1
		}

		if i == len(input)-1 {
			result = append(result, input[len(input)-1:])
		}
	}

	return result
}

func getCoordinates(instruction []string) []int {
	coordinates := []int{0, 0}

	for _, direction := range instruction {
		if direction == "e" {
			coordinates[0] = coordinates[0] + 2
		}

		if direction == "w" {
			coordinates[0] = coordinates[0] - 2
		}

		if direction == "se" {
			coordinates[0] = coordinates[0] + 1

			if isOdd(coordinates[0]) {
				coordinates[1] = coordinates[1] + 1
			}
		}

		if direction == "ne" {
			coordinates[0] = coordinates[0] + 1

			if isEven(coordinates[0]) {
				coordinates[1] = coordinates[1] - 1
			}
		}

		if direction == "sw" {
			coordinates[0] = coordinates[0] - 1

			if isOdd(coordinates[0]) {
				coordinates[1] = coordinates[1] + 1
			}
		}

		if direction == "nw" {
			coordinates[0] = coordinates[0] - 1

			if isEven(coordinates[0]) {
				coordinates[1] = coordinates[1] - 1
			}
		}
	}

	return coordinates
}

func isEven(num int) bool {
	return num%2 == 0
}

func isOdd(num int) bool {
	return !isEven(num)
}
