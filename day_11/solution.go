package main

import (
	"aoc-2020/utils"
	"fmt"
	"strings"
)

const occupied = "#"
const floor = "."
const empty = "L"

func main() {
	data := utils.ReadTextFile("./data.txt")
	var seatMap = make([][]string, len(data))

	for index, row := range data {
		seatMap[index] = strings.Split(row, "")
	}

	fmt.Println(getOccupiedSeatCount(getStabilizedSeatMap(seatMap)))
}

func getStabilizedSeatMap(seatMap [][]string) [][]string {
	newSeats, hasChanged := populateSeats(seatMap)

	if !hasChanged {
		return newSeats
	}

	return getStabilizedSeatMap(newSeats)
}

func getOccupiedSeatCount(seatMap [][]string) int {
	count := 0

	for _, row := range seatMap {
		for _, seat := range row {
			if seat == occupied {
				count++
			}
		}
	}

	return count
}

func populateSeats(seatMap [][]string) ([][]string, bool) {
	var newSeatMap [][]string
	seatStateChanged := false

	for rowIndex, row := range seatMap {
		var seatRow []string

		for seatIndex, seat := range row {
			seatState := getNewSeatState(seatIndex, rowIndex, seatMap)
			seatRow = append(seatRow, seatState)
			if seat != seatState {
				seatStateChanged = true
			}
		}

		newSeatMap = append(newSeatMap, seatRow)
	}

	return newSeatMap, seatStateChanged
}

func getNewSeatState(xIndex int, yIndex int, seatMap [][]string) string {
	adjacentSeats := getAdjacentSeats(xIndex, yIndex, seatMap)
	seat := seatMap[yIndex][xIndex]

	adjacentsEmpty := utils.ContainsString(adjacentSeats, occupied) == false

	if seat == empty && adjacentsEmpty {
		return occupied
	} else if seat == occupied && minimalEqual(adjacentSeats, occupied, 4) {
		return empty
	}

	return seat
}

func minimalEqual(arr []string, value string, min int) bool {
	timesEqual := 0

	for _, item := range arr {
		if item == value {
			timesEqual++
		}
	}

	return timesEqual >= min
}

func getAdjacentSeats(xIndex int, yIndex int, seatMap [][]string) []string {
	xLimit := len(seatMap[0]) - 1
	yLimit := len(seatMap) - 1

	xNext := xIndex + 1
	xPrev := xIndex - 1
	yNext := yIndex + 1
	yPrev := yIndex - 1

	var seats []string

	if xNext <= xLimit {
		seats = append(seats, seatMap[yIndex][xNext])
	}

	if xNext <= xLimit && yPrev >= 0 {
		seats = append(seats, seatMap[yPrev][xNext])
	}

	if xNext <= xLimit && yNext <= yLimit {
		seats = append(seats, seatMap[yNext][xNext])
	}

	if xPrev >= 0 {
		seats = append(seats, seatMap[yIndex][xPrev])
	}

	if xPrev >= 0 && yPrev >= 0 {
		seats = append(seats, seatMap[yPrev][xPrev])
	}

	if xPrev >= 0 && yNext <= yLimit {
		seats = append(seats, seatMap[yNext][xPrev])
	}

	if yNext <= yLimit {
		seats = append(seats, seatMap[yNext][xIndex])
	}

	if yPrev >= 0 {
		seats = append(seats, seatMap[yPrev][xIndex])
	}

	return seats
}
