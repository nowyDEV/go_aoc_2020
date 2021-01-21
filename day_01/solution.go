package main

import (
	"aoc-2020/utils"
	"fmt"
	"sort"
)

func main() {
	data := utils.ReadTextFile("./data.txt")

	var numbers []int

	for _, row := range data {
		number := utils.GetIntFromString(row)

		if isValidValue(number) {
			numbers = append(numbers, utils.GetIntFromString(row))
		}

	}

	sort.Ints(numbers)

	for i := 0; i < len(numbers)-2; i++ {
		for j := len(numbers) - 1; j > 1; j-- {
			a := numbers[i]
			b := numbers[j]

			sum := a + b

			if sum < 2020 {
				if utils.ContainsInt(numbers, 2020-sum) {
					fmt.Println("part 2: ", a*b*(2020-sum))
				}
				continue
			} else if sum == 2020 {
				fmt.Println("part 1: ", a*b)
			}
		}
	}
}

func isValidValue(value int) bool {
	return value < 2021
}
