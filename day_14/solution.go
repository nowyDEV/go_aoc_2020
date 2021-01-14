package main

import (
	"aoc-2020/utils"
	"fmt"
	"strings"
)

type memAddress struct {
	address int
	value   int
}

func main() {
	data := utils.ReadTextFile("./data_test.txt")

	x := 101

	fmt.Printf("%d\t%b\n", x, x)

	fmt.Println(data)

	parsedData := parseInput(data)

	// var addresses []memAddress

	for _, row := range parsedData {
		fmt.Println(row)
	}
}

func parseInput(input []string) (result [][]string) {
	index := -1

	for _, row := range input {
		if isMask(row) {
			index++
			result = append(result, []string{getMask(row)})
		} else {
			result[index] = append(result[index], row)
		}
	}

	return result
}

func isMask(input string) bool {
	return strings.Contains(input, "mask")
}

func getMask(input string) string {
	return strings.Replace(input, "mask = ", "", 1)
}
