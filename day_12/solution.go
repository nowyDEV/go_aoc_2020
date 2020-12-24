package main

import (
	"aoc-2020/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type position struct {
	code  string
	value int
}

type shipPosition struct {
	horizontal position
	vertical   position
	turn       string
}

var directionsR = []string{"N", "E", "S", "W"}
var directionsL = []string{"N", "W", "S", "E"}

var moveCodes = append(directionsR, "F")
var turnCodes = []string{"R", "L"}

func main() {
	data := utils.ReadTextFile("./data.txt")

	shipPos := shipPosition{
		position{
			"E",
			0,
		},
		position{
			"N",
			0,
		},
		"E",
	}

	for _, item := range data {
		code, _ := getValues(item)

		if utils.ContainsString(moveCodes, code) {
			shipPos = handleMove(shipPos, item)
		} else {
			shipPos.turn = handleTurn(shipPos.turn, item)
		}
	}

	fmt.Println(shipPos)
	fmt.Println(shipPos.vertical.value + shipPos.horizontal.value)
}

func handleTurn(currTurn string, turn string) string {
	turnCode, turnDegrees := getValues(turn)
	turnValue := turnDegrees / 90

	if turnCode == "R" {
		currIndex := utils.FindIndex(directionsR, currTurn)
		nextIndex := (currIndex + turnValue) % 4

		return directionsR[nextIndex]
	}

	if turnCode == "L" {
		currIndex := utils.FindIndex(directionsL, currTurn)
		nextIndex := (currIndex + turnValue) % 4

		return directionsL[nextIndex]
	}

	return currTurn
}

func handleMove(shipPosition shipPosition, move string) shipPosition {
	moveCode, moveValue := getValues(move)

	if moveCode == shipPosition.vertical.code {
		shipPosition.vertical.value = shipPosition.vertical.value + moveValue
		return shipPosition
	}

	if moveCode == shipPosition.horizontal.code {
		shipPosition.horizontal.value = shipPosition.horizontal.value + moveValue
		return shipPosition
	}

	if moveCode == "E" || moveCode == "W" {
		shipPosition.horizontal.value = shipPosition.horizontal.value - moveValue

		if shipPosition.horizontal.value >= 0 {
			return shipPosition
		}

		if shipPosition.horizontal.code == "W" {
			shipPosition.horizontal.code = "E"
		} else {
			shipPosition.horizontal.code = "W"
		}
		shipPosition.horizontal.value = shipPosition.horizontal.value * -1

		return shipPosition
	}

	if moveCode == "N" || moveCode == "S" {
		shipPosition.vertical.value = shipPosition.vertical.value - moveValue

		if shipPosition.vertical.value >= 0 {
			return shipPosition
		}

		if shipPosition.vertical.code == "N" {
			shipPosition.vertical.code = "S"
		} else {
			shipPosition.vertical.code = "N"
		}
		shipPosition.vertical.value = shipPosition.vertical.value * -1

		return shipPosition
	}

	if moveCode == "F" {
		concatenated := fmt.Sprintf("%s%d", shipPosition.turn, moveValue)

		return handleMove(shipPosition, concatenated)
	}

	return shipPosition
}

func getValues(input string) (code string, value int) {
	var reValues = regexp.MustCompile("([A-Z]+)([0-9]+)")
	matches := reValues.FindStringSubmatch(input)

	value, err := strconv.Atoi(matches[2])

	if err != nil {
		log.Fatalf("Failed to extract value")
	}

	return matches[1], value
}
