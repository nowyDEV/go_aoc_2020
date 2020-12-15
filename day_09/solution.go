package main

import (
	"aoc-2020/utils"
	"fmt"
	"log"
	"sort"
)

const numOfPreamble = 25

func main() {
	data := utils.ReadTextFile("./data.txt")

	numbers := make([]int, len(data))
	for index, item := range data {
		numbers[index] = utils.GetIntFromString(item)
	}

	incorrectNumber := getIncorrectNumber(numbers, numOfPreamble)
	fmt.Println("solution 1: ", incorrectNumber)

	sumParts := findContiguousThatSumUp(0, 1, numbers, incorrectNumber)
	sort.Ints(sumParts)
	fmt.Println("solution 2: ", sumParts[0]+sumParts[len(sumParts)-1])
}

func findContiguousThatSumUp(startIndex int, endIndex int, numbers []int, sum int) []int {
	checkSum := utils.Sum(numbers[startIndex:endIndex])

	if sum-checkSum == 0 {
		return numbers[startIndex:endIndex]
	}

	if sum-checkSum > 0 {
		return findContiguousThatSumUp(startIndex, endIndex+1, numbers, sum)
	}

	return findContiguousThatSumUp(startIndex+1, endIndex, numbers, sum)
}

func getIncorrectNumber(numbers []int, numOfPreamble int) int {
	for i := numOfPreamble; i < len(numbers); i++ {
		if isSumOfTwo(numbers[i], numbers[i-numOfPreamble:i]) {
			continue
		}
		return numbers[i]
	}

	log.Fatalf("number not found, returning 0")
	return 0
}

func isSumOfTwo(sum int, preambles []int) bool {
	length := len(preambles)

	for i := (length - 1); i >= 0; i-- {
		if preambles[i] < sum && utils.ContainsInt(preambles[0:i], sum-preambles[i]) {
			return true
		}
	}
	return false
}
