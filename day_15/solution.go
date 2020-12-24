package main

import (
	"aoc-2020/utils"
	"fmt"
	"strings"
)

func main() {
	data := utils.ReadTextFile("./data.txt")

	var numbers []int

	for _, numStr := range strings.Split(data[0], ",") {
		numbers = append(numbers, utils.GetIntFromString(numStr))
	}

	fmt.Println(getNthNumber(numbers, 2020-len(numbers)))
	// fmt.Println(getNthNumber(numbers, 30000000-len(numbers)))
}

func getNthNumber(startList []int, limit int) int {
	turns := 0

	for turns < limit {
		startList = append(startList, getNextNumber(startList))
		fmt.Println(getNextNumber(startList))

		turns++
	}

	return startList[len(startList)-1]
}

func getNextNumber(numbers []int) int {
	indexes := findIndexes(numbers, numbers[len(numbers)-1])

	if len(indexes) <= 1 {
		return 0
	}

	return indexes[0] - indexes[1]
}

func findIndexes(list []int, value int) []int {
	var result []int

	for i := (len(list) - 1); i >= 0; i-- {
		if list[i] == value {
			result = append(result, i)
		}

		if len(result) == 2 {
			return result
		}
	}

	return result
}
