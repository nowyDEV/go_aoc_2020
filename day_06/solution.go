package main

import (
	"aoc-2020/utils"
	"fmt"
	"strings"
)

type group struct {
	numOfPeople int
	answers     string
}

type mapOfGroups = map[int]*group

func main() {
	formData := utils.ReadTextFile("../day_06/data.txt")

	fmt.Println(sumAllAnswers(mapAllGroups(formData)))
	fmt.Println(sumCommonAnswers(mapAllGroups(formData)))
}

func mapAllGroups(formData []string) mapOfGroups {
	var mappedGroups mapOfGroups
	mappedGroups = make(mapOfGroups)

	var indexKey int = 0

	mappedGroups[indexKey] = &group{
		0,
		"",
	}

	for _, line := range formData {
		if line == "" {
			indexKey = indexKey + 1
			mappedGroups[indexKey] = &group{
				0,
				"",
			}
		} else {
			mappedGroups[indexKey].answers = mappedGroups[indexKey].answers + line
			mappedGroups[indexKey].numOfPeople = mappedGroups[indexKey].numOfPeople + 1
		}
	}

	return mappedGroups
}

func sumCommonAnswers(mappedGroups mapOfGroups) int {
	sum := 0

	for _, group := range mappedGroups {
		uniqueAnswers := getUniqueChars(group.answers)

		for _, answer := range uniqueAnswers {
			if strings.Count(group.answers, answer) == group.numOfPeople {
				sum = sum + 1
			}
		}
	}

	return sum
}

func sumAllAnswers(mappedGroups mapOfGroups) int {
	sum := 0

	for _, group := range mappedGroups {
		sum = sum + len(getUniqueChars(group.answers))
	}

	return sum
}

func getUniqueChars(input string) []string {
	var charMap map[string]bool
	charMap = make(map[string]bool)

	for _, char := range input {
		charMap[string(char)] = true
	}

	uniqueChars := make([]string, len(charMap))

	i := 0
	for char := range charMap {
		uniqueChars[i] = char
		i++
	}

	return uniqueChars
}
