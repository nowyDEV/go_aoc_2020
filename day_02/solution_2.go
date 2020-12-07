package main

import (
	"aoc-2020/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	text := utils.ReadTextFile("./data.txt")

	type PasswordConfig struct {
		positions    []int
		requiredChar string
		password     string
	}

	var validItems int

	for _, item := range text {
		arr := strings.Split(item, " ")

		config := PasswordConfig{convertListOfStrToInt(strings.Split(arr[0], "-")), strings.Replace(arr[1], ":", "", 1), arr[2]}

		positionsOfRequiredChar := getCharPositions(config.password, config.requiredChar)
		numOfValidPositions := getNumOfOccurences(positionsOfRequiredChar, config.positions)

		if numOfValidPositions == 1 {
			validItems = validItems + 1
		}
	}

	fmt.Println(validItems)
}

func getCharPositions(text string, char string) []int {
	textArr := strings.Split(text, "")
	var positions []int

	for i, item := range textArr {
		if item == char {
			positions = append(positions, i+1)
		}
	}

	return positions
}

func getNumOfOccurences(input []int, list []int) int {
	var occurences int

	for _, item := range input {
		if item == list[0] || item == list[1] {
			occurences = occurences + 1
		}
	}

	return occurences
}

func convertListOfStrToInt(list []string) []int {
	var intList []int

	for _, item := range list {
		integer, err := strconv.Atoi(item)

		if err != nil {
			log.Fatalf("failed to convert")
		}

		intList = append(intList, integer)
	}

	return intList
}
