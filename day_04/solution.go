package main

import (
	"fmt"
	"aoc-2020/utils"
	"strings"
	"regexp"
)

var requiredCodes = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
var numOfRequiredValidCodes = len(requiredCodes)

func main() {
	text := utils.ReadTextFile("../day_04/data.txt")

	fmt.Println(calculateValidPassports(text, 0, 0, 0))
}

func calculateValidPassports(textArr []string, index int, validCodes int, count int) int {
	if (index == (len(textArr) - 1)) {
		valid := validCodes + calculateValidCodes(textArr[index])

		if (valid == numOfRequiredValidCodes) {
			return count + 1
		}

		return count
	}

	if (textArr[index] == "") {
			if (validCodes == numOfRequiredValidCodes) {
				return calculateValidPassports(textArr, index + 1, 0, count + 1)
			} 
			return calculateValidPassports(textArr, index + 1, 0, count)
	}

	valid := validCodes + calculateValidCodes(textArr[index])
	return calculateValidPassports(textArr, index + 1, valid, count)
}

func calculateValidCodes(textRow string) int {
	codeKeyValuePairs := strings.Split(textRow, " ")
	var validCodes = 0

	for _, codeKeyValue := range codeKeyValuePairs {
		codeKey := strings.Split(codeKeyValue, ":")[0]
		codeValue := strings.Split(codeKeyValue, ":")[1]

		isValidCodeKey, isValidCodeValue := isValidCode(codeKey, codeValue)

		// Part 1
		// if (isValidCodeKey) {
		// 	validCodes  = validCodes + 1
		// }

		if (isValidCodeKey && isValidCodeValue) {
			validCodes  = validCodes + 1
		}
	}

	return validCodes
}

func isValidCode(key string, value string) (bool, bool) {
	switch key {
		case "byr":
			return true, utils.IsInRange(utils.GetIntFromString(value), 1920, 2002)
		case "iyr":
			return true, utils.IsInRange(utils.GetIntFromString(value), 2010, 2020)
		case "eyr":
			return true, utils.IsInRange(utils.GetIntFromString(value), 2020, 2030)
		case "hgt":
			return true, validateHeight(value)
		case "hcl":
			return true, utils.IsHexValue(value)
		case "ecl":
			return true, validateEyeColor(value)
		case "pid":
			return true, validatePassportID(value)
		default:
			return false, false
	}
}

var validEyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func validateEyeColor(input string) bool {
	return utils.ContainsString(validEyeColors, input) 
}

func validateHeight(input string) bool {
	matchCm, _ := regexp.MatchString("([0-9]+)cm", input)
	matchInches, _ := regexp.MatchString("([0-9]+)in", input)

	if (matchCm) {
		return utils.IsInRange(utils.GetIntFromString(input), 150, 193)
	}

	if (matchInches) {
		return utils.IsInRange(utils.GetIntFromString(input), 59, 76)
	}

	return false
}

func validatePassportID(input string) bool {
	r, _ := regexp.Compile("([0-9]+)")

	indexes := r.FindStringIndex(input)
	return len(input) == 9 && indexes[0] == 0 && indexes[1] == 9
}