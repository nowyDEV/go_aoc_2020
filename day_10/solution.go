package main

import (
	"aoc-2020/utils"
	"fmt"
	"math"
	"sort"
)

func main() {
	data := utils.ReadTextFile("./data_test.txt")

	numbers := make([]int, len(data))
	for index, item := range data {
		numbers[index] = utils.GetIntFromString(item)
	}

	voltageArr := append([]int{0}, numbers...)
	sort.Ints(voltageArr)

	fmt.Println(voltageArr)

	fmt.Println(multiplyJoltDiffs(voltageArr))
	fmt.Println(getArrangements(voltageArr))
}

func getArrangements(voltageArr []int) int {
	var danks []int
	numOfIntermediates := 0

	for i := 0; i < len(voltageArr)-2; i++ {
		diff := int(math.Abs(float64(voltageArr[i] - voltageArr[i+2])))

		if diff <= 3 {
			numOfIntermediates++
			danks = append(danks, voltageArr[i+1])
		}
	}

	fmt.Println(numOfIntermediates, danks)

	return int(math.Pow(2, float64(numOfIntermediates)))
}

func multiplyJoltDiffs(voltageArr []int) int {
	var numOfThrees = 0
	var numOfOnes = 0

	for i := 0; i < len(voltageArr)-1; i++ {
		diff := voltageArr[i+1] - voltageArr[i]

		if diff == 1 {
			numOfOnes++
		} else if diff == 3 {
			numOfThrees++
		}
	}

	numOfThrees++

	return numOfOnes * numOfThrees
}
