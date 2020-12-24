package main

import (
	"aoc-2020/utils"
	"fmt"
	"strings"
)

func main() {
	data := utils.ReadTextFile("./data.txt")

	busList := getValidBusList(data[1])
	arrival := utils.GetIntFromString(data[0])

	busID, waitTime := getNearestBus(busList, arrival)

	fmt.Println(busID * waitTime)
}

func getNearestBus(busList []int, arrival int) (int, int) {
	nearestTime := arrival
	var nearestBus int

	for i := 0; i < len(busList); i++ {
		mod := arrival % busList[i]

		if mod == 0 {
			return busList[i], 0
		}

		busArrival := busList[i] - mod
		if busArrival < nearestTime {
			nearestTime = busArrival
			nearestBus = busList[i]
		}
	}

	return nearestTime, nearestBus

}

func getValidBusList(input string) []int {
	busses := utils.FilterOut(strings.Split(input, ","), "x")

	var result []int

	for _, bus := range busses {
		result = append(result, utils.GetIntFromString(bus))
	}

	return result
}
