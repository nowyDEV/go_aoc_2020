package main

import (
	"fmt"
	"aoc-2020/utils"
)

const numOfRows = 128
const numOfCols = 8

func main() {
	boardingPasses := utils.ReadTextFile("../day_05/data.txt")

	highestSeatID := 0

	rows := generateIndexArray(numOfRows)
	cols := generateIndexArray(numOfCols)

	for _, pass := range boardingPasses {
		row := getRow(rows, pass[0:7], 0)[0]
		column := getColumn(cols, pass[7:10], 0)[0]

		result := row * 8 + column

		fmt.Println("column", column, pass[7:10])

		if (result > highestSeatID) {
			highestSeatID = result
		}
	}

	fmt.Println(highestSeatID)
}

func getColumn(cols []int, directions string, currentIndex int) []int {
	direction := string(directions[currentIndex])

	if (currentIndex == (len(directions) - 1)) {
		if (direction == "L") {
			return []int{cols[0]}
		}
	
		if (direction == "R") {
			return []int{cols[1]}
		}
	}

	if (direction == "L") {
		return getColumn(cols[0:(len(cols) / 2):len(cols)], directions, currentIndex + 1)
	}

	if (direction == "R") {
		return getColumn(cols[len(cols) / 2:len(cols):len(cols)], directions, currentIndex + 1)
	}

	return []int{0}
}

func getRow(rows []int, directions string, currentIndex int) []int {
	direction := string(directions[currentIndex])

	if (currentIndex == (len(directions) - 1)) {
		if (direction == "F") {
			return []int{rows[0]}
		}
	
		if (direction == "B") {
			return []int{rows[1]}
		}
	}

	if (direction == "F") {
		return getRow(rows[0:(len(rows) / 2):len(rows)], directions, currentIndex + 1)
	}

	if (direction == "B") {
		return getRow(rows[len(rows) / 2:len(rows):len(rows)], directions, currentIndex + 1)
	}

	return []int{0}
}

func generateIndexArray(numOfItems int) []int {
	items := make([]int, numOfItems)
	for i := range items {
		items[i] = i
	}
	return items
}