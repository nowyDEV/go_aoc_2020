package main

import (
	"fmt"
	"aoc-2020/utils"
	"sort"
)

const numOfRows = 128
const numOfCols = 8

func main() {
	boardingPasses := utils.ReadTextFile("../day_05/data.txt")

	rows := generateIndexArray(numOfRows)
	cols := generateIndexArray(numOfCols)

	seatIDs := []int{}

	for _, pass := range boardingPasses {
		seatIDs = append(seatIDs, getSeatID(rows, cols, pass))
	}

	fmt.Println(getHighestSeatID(rows, cols, boardingPasses))
	fmt.Println(getFreeSeat(seatIDs))
}

func getFreeSeat(seatIDs []int) int {
	sort.Ints(seatIDs)

	for i, seatID := range seatIDs {
		if (seatIDs[(i + 1) % len(seatIDs)]  - seatID > 1) {
			return seatID + 1
		}
	}

	return -1
}

func getHighestSeatID(rows []int, cols []int, boardingPasses []string) int {
	highestSeatID := 0

	for _, pass := range boardingPasses {
		seatID := getSeatID(rows, cols, pass)

		if (seatID > highestSeatID) {
			highestSeatID = seatID
		}
	}

	return highestSeatID
}

func getSeatID(rows []int, cols []int, boardingPass string) int {
	row := getRow(rows, boardingPass[0:7], 0)[0]
	column := getColumn(cols, boardingPass[7:10], 0)[0]

	return row * 8 + column
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