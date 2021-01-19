package main

import (
	"aoc-2020/utils"
	"strings"
	"testing"
)

// BenchmarkBigLen Test it
func BenchmarkBigLen(b *testing.B) {
	data := utils.ReadTextFile("./data.txt")

	var numbers []int

	for _, numStr := range strings.Split(data[0], ",") {
		numbers = append(numbers, utils.GetIntFromString(numStr))
	}

	getNthNumber(numbers, 30000-len(numbers))
}
