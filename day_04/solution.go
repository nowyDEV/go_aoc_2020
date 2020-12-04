package main

import (
	"fmt"
	"utils/readfile"
	"strings"
)

var requiredCodes = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
var numOfRequiredValidCodes = len(requiredCodes)
var requiredCodesString = strings.Join(requiredCodes[:], ":") + ":"

func main() {
	text := readfile.GetFileContents("../day_04/data.txt")

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
	codeValuePairs := strings.Split(textRow, " ")
	var validCodes = 0

	for _, codeValuePair := range codeValuePairs {
		code := strings.Split(codeValuePair, ":")[0]

		if (isValid(code)) {
			validCodes  = validCodes + 1
		}
	}

	return validCodes
}

func isValid(input string) bool {
	return strings.Contains(requiredCodesString, input)
}