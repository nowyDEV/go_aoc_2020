package main

import (
	"aoc-2020/utils"
	"fmt"
	"strings"
)

func main() {
	data := utils.ReadTextFile("./data_test.txt")

	cups := parseData(data[0])

	fmt.Println("startCups: ", cups)

	fmt.Println(playGame(cups, 9))
}

func parseData(input string) (result []int) {
	numStrings := strings.Split(input, "")

	for _, str := range numStrings {
		result = append(result, utils.GetIntFromString(str))
	}

	return result
}

func playGame(cups []int, maxTurns int) []int {
	max := len(cups) - 1
	turns := 0
	currIndex := 0

	for turns < maxTurns {
		currentCup := cups[currIndex]
		fmt.Println("currentCup", currentCup)

		picked := make([]int, 3)
		copy(picked, getNextItems(cups, currIndex, 3))

		fmt.Println("picked", picked)

		restCups := utils.SubtractSlice(cups, picked)

		fmt.Println("restCups", restCups)

		destinationCup := pickDestinationCup(cups[currIndex]-1, picked, utils.SubtractSlice(restCups, []int{currentCup}))

		fmt.Println("destinationCup", destinationCup)

		destinationIndex := utils.FindIndexNums(restCups, destinationCup)
		currentCupIndex := utils.FindIndexNums(restCups, currentCup)

		cups = concatSlices(restCups[:destinationIndex+1], picked, restCups[destinationIndex+1:])

		if destinationIndex < currentCupIndex {
			cups = moveSlice(cups, 3)
		}

		fmt.Println("newCups", cups)
		fmt.Println("--------------------------------------------")

		if currIndex < max {
			currIndex++
		} else {
			currIndex = 0
		}

		turns++
	}

	return cups
}

func getNextItems(list []int, index int, numOfItems int) (result []int) {
	len := len(list)

	for i := 1; i < 4; i++ {
		nextIndex := (index + i) % len
		result = append(result, list[nextIndex])
	}

	return result
}

func moveSlice(list []int, move int) []int {
	len := len(list)

	result := make([]int, len)

	for index, item := range list {
		nextIndex := index - move

		if nextIndex < 0 {
			nextIndex = len + nextIndex
		}

		result[nextIndex] = item
	}

	return result
}

func concatSlices(slices ...[]int) (result []int) {
	for i := 0; i < len(slices); i++ {
		result = append(result, slices[i]...)
	}

	return result
}

func pickDestinationCup(startValue int, pickedCups []int, restCups []int) int {
	destinationCup := startValue

	for {
		if utils.ContainsInt(pickedCups, destinationCup) {
			destinationCup--
			continue
		}

		if utils.ContainsInt(restCups, destinationCup) {
			return destinationCup
		}

		lowestLabel := utils.GetLowestNumber(restCups)
		if destinationCup < lowestLabel {
			return utils.GetHighestNumber(restCups)
		}
	}
}
