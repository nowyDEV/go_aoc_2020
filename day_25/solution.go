package main

import (
	"aoc-2020/utils"
	"fmt"
)

func main() {
	data := utils.ReadTextFile("./data.txt")

	cardKey, doorKey := parseInput(data)
	cardLoopSize := getLoopSize(cardKey)
	doorLoopSize := getLoopSize(doorKey)

	fmt.Println(getEncryptionKey(doorKey, cardLoopSize), getEncryptionKey(cardKey, doorLoopSize))
}

func parseInput(input []string) (cardKey, doorKey int) {
	return utils.GetIntFromString(input[0]), utils.GetIntFromString(input[1])
}

const remainder = 20201227
const multiplier = 7

func getLoopSize(key int) int {
	loopSize := 0
	value := 1

	for value != key {
		value = (value * multiplier) % remainder
		loopSize++
	}

	return loopSize
}

func getEncryptionKey(key int, loopSize int) int {
	result := 1
	for i := 0; i < loopSize; i++ {
		result = (result * key) % remainder
	}

	return result
}
