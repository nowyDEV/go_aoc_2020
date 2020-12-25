package main

import (
	"aoc-2020/utils"
	"fmt"
	"strings"
)

type field struct {
	name        string
	validValues []string
}

func main() {
	data := utils.ReadTextFile("./data.txt")

	fmt.Println(getScanningErrorRate(data))
}

func getScanningErrorRate(data []string) int {
	fields, _, nearbyTickets := parseInput(data)

	var invalidValues []int
	var sum int

	for _, ticket := range nearbyTickets {
		invalidValues = append(invalidValues, getInvalidTicketValues(fields, ticket)...)
	}

	for _, value := range invalidValues {
		sum += value
	}

	return sum
}

func getInvalidTicketValues(fields []field, ticket []int) []int {
	var invalidValues []int

	for _, ticketValue := range ticket {
		isValid := validateValue(fields, ticketValue)

		if isValid == false {
			invalidValues = append(invalidValues, ticketValue)
		}
	}

	return invalidValues
}

func validateValue(fields []field, value int) bool {
	for _, fieldT := range fields {
		for _, validValue := range fieldT.validValues {
			ranges := strings.Split(validValue, "-")
			isValid := utils.IsInRange(value, utils.GetIntFromString(ranges[0]), utils.GetIntFromString(ranges[1]))
			if isValid == true {
				return true
			}
		}
	}

	return false
}

func parseInput(input []string) ([]field, []int, [][]int) {
	var fields []field
	var yourTicket []int
	var nearbyTickets [][]int

	for index, row := range input {
		if strings.Contains(row, "or") {
			fields = append(fields, parseField(row))
		}

		if row == "your ticket:" {
			yourTicket = append(yourTicket, parseTicketValues(input[index+1])...)
		}

		if row == "nearby tickets:" {
			tickets := input[index+1:]

			for _, ticket := range tickets {
				nearbyTickets = append(nearbyTickets, parseTicketValues(ticket))
			}
		}
	}

	return fields, yourTicket, nearbyTickets
}

func parseField(input string) field {
	var result field
	parts := strings.Split(input, " ")

	result.name = strings.Replace(parts[0], ":", "", 1)
	for index, part := range parts {
		if part == "or" {
			result.validValues = append(result.validValues, parts[index-1], parts[index+1])
		}
	}

	return result
}

func parseTicketValues(input string) []int {
	var result []int
	values := strings.Split(input, ",")

	for _, value := range values {
		result = append(result, utils.GetIntFromString(value))
	}

	return result
}
