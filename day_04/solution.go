package main

import (
	"fmt"
	"utils/readfile"
	"strings"
	"strconv"
	"encoding/hex"
	"regexp"
	"log"
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
	codeKeyValuePairs := strings.Split(textRow, " ")
	var validCodes = 0

	for _, codeKeyValue := range codeKeyValuePairs {
		codeKey := strings.Split(codeKeyValue, ":")[0]
		codeValue := strings.Split(codeKeyValue, ":")[1]

		// Part 1
		// if (isValidCodeKey(codeKey)) {
		// 	validCodes  = validCodes + 1
		// }

		if (isValidCodeKey(codeKey) && isValidCodeValue(codeKey, codeValue)) {
			validCodes  = validCodes + 1
		}
	}

	return validCodes
}

func isValidCodeValue(key string, value string) bool {
	switch key {
		case "byr":
			return isInRange(getIntFromString(value), 1920, 2002)
		case "iyr":
			return isInRange(getIntFromString(value), 2010, 2020)
		case "eyr":
			return isInRange(getIntFromString(value), 2020, 2030)
		case "hgt":
			return validateHeight(value)
		case "hcl":
			return validateHex(value)
		case "ecl":
			return validateEyeColor(value)
		case "pid":
			return validatePassportID(value)
		default:
			return false
	}
}

func isValidCodeKey(input string) bool {
	return strings.Contains(requiredCodesString, input)
}

func isInRange(input int, start int, end int) bool {
	return input >= start && input <= end
}

func validateHex(input string) bool {
	_, err := hex.DecodeString(strings.Replace(input, "#", "", 1))
	if (err != nil) {
		return false
	}
	return len(input) == 7
}

var eyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func validateEyeColor(input string) bool {
	return contains(eyeColors, input) 
}

func validateHeight(input string) bool {
	matchCm, _ := regexp.MatchString("([0-9]+)cm", input)
	matchInches, _ := regexp.MatchString("([0-9]+)in", input)

	if (matchCm) {
		return isInRange(getIntFromString(input), 150, 193)
	}

	if (matchInches) {
		return isInRange(getIntFromString(input), 59, 76)
	}

	return false
}

func getIntFromString(input string) int {
	r, _ := regexp.Compile("([0-9]+)")
	value, err := strconv.Atoi(r.FindString(input))

	if (err != nil) {
		log.Fatalf("failed to convert")
	}

	return value
}

func contains(s []string, e string) bool {
	for _, a := range s {
			if a == e {
					return true
			}
	}
	return false
}

func validatePassportID(input string) bool {
	r, _ := regexp.Compile("([0-9]+)")

	indexes := r.FindStringIndex(input)
	return len(input) == 9 && indexes[0] == 0 && indexes[1] == 9
}