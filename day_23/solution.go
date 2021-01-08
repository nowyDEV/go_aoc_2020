package main

import (
	"aoc-2020/utils"
	"fmt"
	"strings"
)

func main() {
	data := utils.ReadTextFile("./data.txt")

	cups := parseData(data[0])

	cupsAfterGame := playGame(cups, 100)
	fmt.Println(joinCups(cupsAfterGame, 1))
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
	currentCup := cups[0]

	for turns < maxTurns {
		currentCupIndex := utils.FindIndexNums(cups, currentCup)

		picked := make([]int, 3)
		copy(picked, getNextItems(cups, currentCupIndex, 3))

		restCups := utils.SubtractSlice(cups, picked)

		destinationCup := pickDestinationCup(cups[currentCupIndex]-1, picked, utils.SubtractSlice(restCups, []int{currentCup}))

		destinationIndex := utils.FindIndexNums(restCups, destinationCup)

		cups = concatSlices(restCups[:destinationIndex+1], picked, restCups[destinationIndex+1:])

		newCurrentCupIndex := utils.FindIndexNums(cups, currentCup)

		cups = moveSlice(cups, newCurrentCupIndex-currentCupIndex)

		currentCupIndex = utils.FindIndexNums(cups, currentCup)

		if currentCupIndex < max {
			currentCup = cups[currentCupIndex+1]
		} else {
			currentCup = cups[0]
		}

		turns++
	}

	return cups
}

func joinCups(list []int, startValue int) (result string) {
	startIndex := utils.FindIndexNums(list, startValue)

	for i := startIndex + 1; i < len(list)+startIndex; i++ {
		result = fmt.Sprintf("%s%d", result, list[i%9])
	}

	return result
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
